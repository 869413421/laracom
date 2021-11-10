<?php

namespace App\Providers;

use App\MicroApi\Services\UserService;
use GuzzleHttp\Client as HttpClient;
use Illuminate\Support\ServiceProvider;

use function foo\func;

// use Laravel\Cashier\Cashier;

class AppServiceProvider extends ServiceProvider
{
    /**
     * Bootstrap any application services.
     *
     * @return void
     */
    public function boot()
    {
        // Cashier::useCurrency(config('cart.currency'), config('cart.currency_symbol'));
    }

    /**
     * Register any application services.
     *
     * @return void
     */
    public function register()
    {
        // ����������ͻ��˵�APP����
        $this->app->singleton('HttpClient',function ($app){
            return new HttpClient(
                [
                    'base_uri' => config('services.micro.api_gateway'),
                    'timeout'  => config('services.micro.timeout'),
                    'headers'  => [
                        'Content-Type' => 'application/json'
                    ]
                ]
            );
        });

        // ������microApiService
        $this->app->singleton('microUserService',function($app){
            return new UserService();
        });
    }
}
