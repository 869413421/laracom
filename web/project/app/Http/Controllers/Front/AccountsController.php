<?php

namespace App\Http\Controllers\Front;

use App\Http\Controllers\Controller;
use App\MicroApi\Services\UserService;
use App\Shop\Couriers\Repositories\Interfaces\CourierRepositoryInterface;
use App\Shop\Customers\Repositories\CustomerRepository;
use App\Shop\Customers\Repositories\Interfaces\CustomerRepositoryInterface;
use App\Shop\Orders\Order;
use App\Shop\Orders\Transformers\OrderTransformable;
use Illuminate\Http\Request;

class AccountsController extends Controller
{
    use OrderTransformable;

    /**
     * @var CustomerRepositoryInterface
     */
    private $customerRepo;

    /**
     * @var CourierRepositoryInterface
     */
    private $courierRepo;

    /**
     * @var $userService
     */
    private $userService;

    /**
     * AccountsController constructor.
     *
     * @param CourierRepositoryInterface  $courierRepository
     * @param CustomerRepositoryInterface $customerRepository
     */
    public function __construct(
        CourierRepositoryInterface $courierRepository,
        CustomerRepositoryInterface $customerRepository,
        UserService $userService
    ) {
        $this->customerRepo = $customerRepository;
        $this->courierRepo  = $courierRepository;
        $this->userService  = $userService;
    }

    public function index()
    {
        $customer = $this->customerRepo->findCustomerById(auth()->user()->id);

        $customerRepo = new CustomerRepository($customer);
        $orders       = $customerRepo->findOrders(['*'], 'created_at');

        $orders->transform(function (Order $order) {
            return $this->transformOrder($order);
        });

        $orders->load('products');

        $addresses = $customerRepo->findAddresses();

        return view('front.accounts',
            [
                'customer' => $customer,
                'orders' => $this->customerRepo->paginateArrayResults($orders->toArray(),
                    15),
                'addresses' => $addresses
            ]);
    }

    public function profile(Request $request)
    {
        dd($request->user());
    }
}
