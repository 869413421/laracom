<?php

use App\MicroApi\Items\ProductItem;
use App\MicroApi\Services\ProductService;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Route;
use Illuminate\Support\Str;

/*
|--------------------------------------------------------------------------
| API Routes
|--------------------------------------------------------------------------
|
| Here is where you can register API routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| is assigned the "api" middleware group. Enjoy building your API!
|
*/

Route::middleware('auth:api')->get('/user', function (Request $request) {
    return $request->user();
});

Route::get("/product/test", function (Request $request) {
    $service = new ProductService();
    $product = new ProductItem();

    $product->brand_id    = 1;
    $product->sku         = Str::random(16);
    $product->name        = mb_convert_encoding("微服务架构", 'UTF-8', 'UTF-8,GBK,GB2312,BIG5');
    $product->slug        = 'microservice';
    $product->description = mb_convert_encoding("基于 Laravel + Go Micro 框架构建微服务系统", 'UTF-8', 'UTF-8,GBK,GB2312,BIG5');
    $product->cover       = 'https://laravel.gstatics.cn/wp-content/uploads/2019/06/94fe5973d09b0ad753082b6b1ba46f3d.jpeg';
    $product->price       = 199;
    $product->sale_price  = 99;
    $product->quantity    = 1000;

    $service->create($product);

    dd($service->getById($product->id));
});
