package service

import (
	"context"
	"go-mall/dao"
	"go-mall/model"
	"go-mall/pkg/e"
	"go-mall/pkg/util"
	"go-mall/serializer"
	"mime/multipart"
	"strconv"
	"sync"
)

type ProductService struct {
	ID            uint   `json:"id" form:"id"`
	Name          string `json:"name" form:"name"`
	CategoryId    uint   `json:"category_id" form:"category_id"`
	Title         string `json:"title" form:"title"`
	Info          string `json:"info" form:"info"`
	ImgPath       string `json:"img_path" form:"img_path"`
	Price         string `json:"price" form:"price"`
	DiscountPrice string `json:"discount_price" form:"discount_price"`
	OnSale        bool   `json:"on_sale" form:"on_sale"`
	Num           int    `json:"num" form:"num"`
	BossID        int    `json:"boss_id" form:"boss_id"`
	BossName      string `json:"boss_name" form:"boss_name"`
	BossAvatar    string `json:"boss_avatar" form:"boss_avatar"`

	// 分页相关的信息
	model.BasePage
}

// Create 创建商品
func (service *ProductService) Create(ctx context.Context, uId uint, files []*multipart.FileHeader) serializer.Response {
	code := e.Success

	var boss *model.User
	var err error

	// dao层对象
	useDao := dao.NewUserDao(ctx)
	productDao := dao.NewProductDao(ctx)

	// 查询当前用户信息
	boss, _ = useDao.GetUserById(uId)

	// 以第一张作为商品封面图
	tmp, _ := files[0].Open()
	path, err := UploadProductToLocalStatic(tmp, uId, service.Name)

	if err != nil {
		code := e.ErrorProductImgUpload
		util.LogrusObj.Info("保存商品图片 err: ", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 持久化商品数据信息
	product := &model.Product{
		Name:          service.Name,
		CategoryID:    service.CategoryId,
		Title:         service.Title,
		Info:          service.Info,
		ImgPath:       path,
		Price:         service.Price,
		DiscountPrice: service.DiscountPrice,
		OnSale:        true,
		Num:           service.Num,
		BossID:        service.BossID,
		BossName:      boss.UserName,
		BossAvatar:    boss.Avatar,
	}
	err = productDao.CreateProduct(product)
	if err != nil {
		code := e.ErrorProductCreate
		util.LogrusObj.Info("持久化商品信息 err: ", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 并发保存多个图片
	wg := new(sync.WaitGroup)
	wg.Add(len(files))
	for index, file := range files {
		num := strconv.Itoa(index)
		productImgDao := dao.NewProductImgDaoByDB(productDao.DB)
		tmp, _ := file.Open()
		// 保存到本地时 加入num区分图片
		path, err = UploadProductToLocalStatic(tmp, uId, service.Name+num)
		if err != nil {
			code := e.ErrorProductImgUpload
			util.LogrusObj.Info("保存商品图片 err: ", err)
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
		// 创建商品图片信息
		productImg := &model.ProductImg{
			ProductID: product.ID,
			ImgPath:   path,
		}
		err = productImgDao.CreateProductImg(productImg)
		if err != nil {
			code := e.ErrorProductImgUpload
			util.LogrusObj.Info("保存商品图片 err: ", err)
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
		wg.Done()
	}
	wg.Wait()

	// 返回
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildProduct(product),
	}
}

// Products 获取商品列表
func (service *ProductService) Products(ctx context.Context) serializer.Response {
	code := e.Success
	var err error
	var products []*model.Product

	// 分页信息
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	// 按照商品分类查找
	condition := make(map[string]interface{})
	if service.CategoryId != 0 {
		condition["category_id"] = service.CategoryId
	}

	// 获取业务层对象
	productDao := dao.NewProductDao(ctx)

	//查询数据
	// 1 获取数据分类总数
	total, err := productDao.CountProductByCondition(condition)
	if err != nil {
		//加入日志记录
		util.LogrusObj.Info("获取商品数据失败 err: ", err)
		code = e.ErrorProductGet
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
			Data:   "获取商品数据失败",
		}
	}

	// 并发处理
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		productDao = dao.NewProductDaoByDB(productDao.DB)
		products, _ = productDao.ListProductsByCondition(condition, service.BasePage)
		wg.Done()
	}()
	wg.Wait()

	// 返回
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildListResponse(serializer.BuildProducts(products), uint(total)),
	}
}

// SearchProduct 搜索商品信息
func (service *ProductService) SearchProduct(ctx context.Context) serializer.Response {
	code := e.Success
	var err error
	var count int64
	var products []*model.Product
	var productDao = dao.NewProductDao(ctx)

	// 分页信息
	if service.PageSize == 0 {
		service.PageSize = 15
	}

	products, count, err = productDao.SearchProduct(service.Info, service.BasePage)
	if err != nil {
		//加入日志记录
		util.LogrusObj.Info("查询商品数据失败 err: ", err)
		code = e.ErrorProductGet
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
			Data:   "查询商品数据失败",
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildListResponse(serializer.BuildProducts(products), uint(count)),
	}
}

// Show 展示商品详细信息和图片
func (service *ProductService) Show(ctx context.Context, id string) serializer.Response {
	code := e.Success
	pId, _ := strconv.Atoi(id)
	var product *model.Product
	var err error
	var productDao = dao.NewProductDao(ctx)

	product, err = productDao.GetProductById(uint(pId))
	if err != nil {
		//加入日志记录
		util.LogrusObj.Info("查询商品数据失败 err: ", err)
		code = e.ErrorProductGet
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
			Data:   "查询商品数据失败",
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildProduct(product),
	}

}
