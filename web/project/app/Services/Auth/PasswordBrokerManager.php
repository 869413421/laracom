<?php
namespace App\Services\Auth;

use Closure;
use Illuminate\Auth\Passwords\PasswordBrokerManager as BasePasswordBrokerManager;

class PasswordBrokerManager extends BasePasswordBrokerManager
{
    /**
     * Create a token repository instance based on the given configuration.
     *
     * @param  array  $config
     * @return \Illuminate\Auth\Passwords\TokenRepositoryInterface
     */
    protected function createTokenRepository(array $config)
    {
        return new ServiceTokenRepository();
    }

    public function sendResetLink(array $credentials, Closure $callback = null)
    {
        dd("link");
        return static::RESET_LINK_SENT;
    }

    public function reset(array $credentials, Closure $callback)
    {
        dd("reset");
    }
}