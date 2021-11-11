<?php

namespace App\Providers;

use App\Services\Auth\JwtGuard;
use App\Services\Auth\MicroUserProvider;
use Illuminate\Foundation\Support\Providers\AuthServiceProvider as ServiceProvider;
use Illuminate\Support\Facades\Auth;

class AuthServiceProvider extends ServiceProvider
{
    /**
     * The policy mappings for the application.
     *
     * @var array
     */
    protected $policies
        = [
            'App\Model' => 'App\Policies\ModelPolicy',
        ];

    /**
     * Register any authentication / authorization services.
     *
     * @return void
     */
    public function boot()
    {
        $this->registerPolicies();

        // ��չ User Provider
        Auth::provider('micro', function($app, array $config) {
            // ����һ��Illuminate\Contracts\Auth\UserProviderʵ��...
            return new MicroUserProvider($config['model']);
        });

        // ��չ Auth Guard
        Auth::extend('jwt', function($app, $name, array $config) {
            // ����һ��Illuminate\Contracts\Auth\Guardʵ��...
            return new JwtGuard(Auth::createUserProvider($config['provider']), $app->make('request'));
        });
    }
}
