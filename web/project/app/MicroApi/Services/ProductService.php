<?php
namespace App\MicroApi\Services;

use App\MicroApi\Exceptions\RpcException;
use App\MicroApi\Facades\HttpClient;
use App\MicroApi\Items\ProductItem;
use Illuminate\Support\Facades\Log;

class ProductService
{
    use DataHandler;

    protected $servicePrefix = '/product/productService';

    /**
     * @param $product
     * @return ProductItem
     * @throws RpcException
     */
    public function create(ProductItem $product)
    {
        $path = $this->servicePrefix . '/create';
        $options = ['json' => $product];
        try {
            $response = HttpClient::post($path, $options);
        } catch (\Exception $exception) {
            Log::error("MicroApi.ProductService.Create Call Failed: " . $exception->getMessage());
            throw new RpcException("����Զ�̷���ʧ��");
        }
        $result = $this->decode($response->getBody()->getContents());
        return isset($result->product) ? $product->fillAttributes($result->product) : null;
    }

    /**
     * ��ȡ��Ʒ����
     * @return \Illuminate\Support\Collection|null
     * @throws RpcException
     */
    public function getAll()
    {
        $path = $this->servicePrefix . '/getAll';
        try {
            $response = HttpClient::get($path);
        } catch (\Exception $exception) {
            Log::error("MicroApi.ProductService.GetAll Call Failed: " . $exception->getMessage());
            throw new RpcException("����Զ�̷���ʧ��");
        }
        $result = $this->decode($response->getBody()->getContents());
        // ������Ʒ����
        return isset($result->products) ? collect($result->products)->map(function ($item) {
            $product = new ProductItem();
            return $product->fillAttributes($item);
        })->reject(function ($item){
            return empty($item);
        }) : null;
    }

    /**
     * ������Ʒ ID ��ȡ��Ʒ��Ϣ
     * @param $id
     * @param $withRelations
     * @return ProductItem|null
     * @throws RpcException
     */
    public function getById($id, $withRelations = false)
    {
        if ($withRelations == false) {
            $path = $this->servicePrefix . '/get';
        } else {
            $path = $this->servicePrefix . '/getDetail';
        }
        $product = new ProductItem();
        $product->id = $id;
        $options = ['json' => $product];
        try {
            $response = HttpClient::post($path, $options);
        } catch (\Exception $exception) {
            Log::error("MicroApi.ProductService.Get Call Failed: " . $exception->getMessage());
            throw new RpcException("����Զ�̷���ʧ��");
        }
        $result = $this->decode($response->getBody()->getContents());
        return isset($result->product) ? $product->fillAttributes($result->product) : null;
    }

    /**
     * ͨ����Ʒ������ȡ��ϸ
     * @param $slug
     * @return ProductItem|null
     * @throws RpcException
     */
    public function getBySlug($slug)
    {
        $path = $this->servicePrefix . '/get';
        $product = new ProductItem();
        $product->slug = $slug;
        $options = ['json' => $product];
        try {
            $response = HttpClient::post($path, $options);
        } catch (\Exception $exception) {
            Log::error("MicroApi.ProductService.Get Call Failed: " . $exception->getMessage());
            throw new RpcException("����Զ�̷���ʧ��");
        }
        $result = $this->decode($response->getBody()->getContents());
        return isset($result->user) ? $product->fillAttributes($product) : null;
    }

    /**
     * ������Ʒ��Ϣ
     * @param ProductItem $product
     * @return ProductItem
     * @throws RpcException
     */
    public function update(ProductItem $product)
    {
        $path = $this->servicePrefix . '/update';
        $options = ['json' => $product];
        try {
            $response = HttpClient::post($path, $options);
        } catch (\Exception $exception) {
            Log::error("MicroApi.ProductService.Update Call Failed: " . $exception->getMessage());
            throw new RpcException("����Զ�̷���ʧ��");
        }
        $result = $this->decode($response->getBody()->getContents());
        return $product->fillAttributes($result->product);
    }

    /**
     * ɾ����Ʒ��Ϣ
     * @param $productId
     * @return bool
     * @throws RpcException
     */
    public function delete($productId)
    {
        $path = $this->servicePrefix . '/delete';
        $options = ['json' => ['id' => $productId]];
        try {
            $response = HttpClient::post($path, $options);
        } catch (\Exception $exception) {
            Log::error("MicroApi.ProductService.Delete Call Failed: " . $exception->getMessage());
            throw new RpcException("����Զ�̷���ʧ��");
        }
        $result = $this->decode($response->getBody()->getContents());
        return empty($result->product) ? true : false;
    }
}