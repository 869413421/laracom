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
        // TODO: Implement sendResetLink() method.
    }

    public function reset(array $credentials, Closure $callback)
    {
        // TODO: Implement reset() method.
    }
}