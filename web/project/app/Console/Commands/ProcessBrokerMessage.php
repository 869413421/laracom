<?php

namespace App\Console\Commands;

use App\MicroApi\Items\UserItem;
use App\MicroApi\Services\UserService;
use App\Services\Broker\BrokerService;
use Illuminate\Console\Command;

class ProcessBrokerMessage extends Command
{
    protected $signature = 'process:broker-message';

    protected $description = 'Subscribe and process message from micro broker';

    /**
     * @var UserService
     */
    protected $userService;

    /**
     * Create a new command instance.
     *
     * @return void
     */
    public function __construct()
    {
        parent::__construct();
        $this->userService = resolve("microUserService");
    }

    /**
     * Execute the console command.
     *
     * @return mixed
     */
    public function handle()
    {
        $broker = new BrokerService();
        $broker->subscribe('password.reset', function ($message) {
            // ������Ϣ����
            $message       = $message->getBody();
            $passwordReset = json_decode(base64_decode($message['Body']));
            $email         = $passwordReset->email;
            $token         = $passwordReset->token;
            // ���������ʼ�
            $user = $this->userService->getByEmail($email);
            if ($user) {
                $model = new UserItem();
                $model->fillAttributes($user);
                $model->sendPasswordResetNotification($token);
                $this->info('���������ʼ��ѷ���[email:' . $email . ']');
            } else {
                $this->error('ָ���û�������[email:' . $email . ']');
            }
        });
        $broker->wait();
    }
}