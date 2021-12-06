package main

import (
	"fmt"
	"github.com/869413421/laracom/product-service/db"
	"github.com/869413421/laracom/product-service/handler"
	"github.com/869413421/laracom/product-service/model"
	pb "github.com/869413421/laracom/product-service/proto/product"
	"github.com/869413421/laracom/product-service/repo"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/util/log"
)

func main() {
	//1.初始化数据库，执行数据迁移
	database, err := db.CreateConnection()
	defer database.Close()
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	// 数据库迁移（商品、图片、品牌、类目、属性相关数据表）
	database.AutoMigrate(&model.Product{})
	database.AutoMigrate(&model.ProductImage{})
	database.AutoMigrate(&model.Brand{})
	database.AutoMigrate(&model.Category{})
	database.AutoMigrate(&model.Attribute{})
	database.AutoMigrate(&model.AttributeValue{})
	database.AutoMigrate(&model.ProductAttribute{})

	// 初始化 Repo 实例用于后续数据库操作
	productRepo := &repo.ProductRepository{Db: database}
	imageRepo := &repo.ImageRepository{Db: database}
	brandRepo := &repo.BrandRepository{Db: database}
	categoryRepo := &repo.CategoryRepository{Db: database}
	attributeRepo := &repo.AttributeRepository{Db: database}

	//3.创建微服务
	srv := micro.NewService(micro.Name("laracom.service.product"), micro.Version("latest"))
	srv.Init()

	//4.注册服务处理器
	pb.RegisterProductServiceHandler(srv.Server(), &handler.ProductService{ProductRepo: productRepo})
	pb.RegisterImageServiceHandler(srv.Server(), &handler.ImageService{ImageRepo: imageRepo})
	pb.RegisterBrandServiceHandler(srv.Server(), &handler.BrandService{BrandRepo: brandRepo})
	pb.RegisterCategoryServiceHandler(srv.Server(), &handler.CategoryService{CategoryRepo: categoryRepo})
	pb.RegisterAttributeServiceHandler(srv.Server(), &handler.AttributeService{AttributeRepo: attributeRepo})

	//5.启动服务
	err = srv.Run()
	if err != nil {
		fmt.Println(err)
	}
}
