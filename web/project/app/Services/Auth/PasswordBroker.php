<?php
namespace App\Services\Auth;

use Closure;
use Illuminate\Auth\Passwords\PasswordBroker as BasePasswordBroker;

class PasswordBroker extends BasePasswordBroker
{
    /**
     * Send a password reset link to a user.
     *
     * @param  array  $credentials
     * @return string
     */
    public function sendResetLink(array $credentials, Closure $callback = null)
    {
        // ����û��Ƿ����
        $user = $this->getUser($credentials);
        if (is_null($user)) {
            return static::INVALID_USER;
        }
        // ���ڵĻ��򴴽���Ӧ���������ü�¼���ʼ����Ͳ����첽ȥ��
        $this->tokens->create($user);
        return static::RESET_LINK_SENT;
    }
}