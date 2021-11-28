<?php

namespace App\Http\Controllers\Front;

use App\Http\Controllers\Controller;
use App\Services\Customer\UserService;
use App\Shop\Couriers\Repositories\Interfaces\CourierRepositoryInterface;
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
        // 用户信息
        $user = auth()->user();

        // 分页订单信息
        $orders = $this->userService->getPaginatedOrdersByUserId($user->id);
        $orders->transform(function (Order $order) {
            return $this->transformOrder($order);
        });

        // 地址信息
        $addresses = $this->userService->getAddressesByUserId($user->id);

        return view('front.accounts', [
            'customer' => $user,
            'orders' => $orders,
            'addresses' => $addresses
        ]);
    }

    public function profile(Request $request)
    {
        dd($request->user());
    }
}
