<?php


namespace App\MicroApi\Items;


use Illuminate\Contracts\Auth\Authenticatable;

class UserItem implements Authenticatable
{
    public $id;
    public $name;
    public $email;
    public $password;
    public $status;

    protected $hidden = ['password'];

    /**
     * fillAttributes JSONÌî³äÎª¶ÔÏó
     *
     * @param $data
     *
     * @return $this
     */
    public function fillAttributes($data)
    {
        if (is_object($data)) {
            $data = get_object_vars($data);
        }

        foreach ($data as $key => $value) {
            if (in_array($key, $this->hidden)) {
                continue;
            }
            switch ($key) {
                case 'id':
                    $this->id = $value;
                    break;
                case 'name':
                    $this->name = $value;
                    break;
                case 'email':
                    $this->email = $value;
                    break;
                case 'status':
                    $this->status = $value;
                    break;
                default:
                    break;
            }
        }
        return $this;
    }

    public function getAuthIdentifierName()
    {
        return 'id';
    }

    public function getAuthIdentifier()
    {
        return $this->id;
    }

    public function getAuthPassword()
    {
        // TODO: Implement getAuthPassword() method.
    }

    public function getRememberToken()
    {
        // TODO: Implement getRememberToken() method.
    }

    public function setRememberToken($value)
    {
        // TODO: Implement setRememberToken() method.
    }

    public function getRememberTokenName()
    {
        // TODO: Implement getRememberTokenName() method.
    }
}