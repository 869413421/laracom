<?php

namespace App\Providers;

use GuzzleHttp\Client as HttpClient;
use Illuminate\Support\ServiceProvider;

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
        // 单例绑定请求客户端到APP容器
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
    }
}
