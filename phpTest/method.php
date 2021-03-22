<?php
/*
    记录一下学习,用到的php系统方法
*/


header("Content-type: text/html; charset=UTF-8");
/*
The Cache-Control is per the HTTP 1.1 spec for clients and proxies (and implicitly required by some clients next to Expires).
The Pragma is per the HTTP 1.0 spec for prehistoric clients. The Expires is per the HTTP 1.0 and 1.1 specs for clients and proxies.
In HTTP 1.1, the Cache-Control takes precedence over Expires, so it's after all for HTTP 1.0 proxies only.
 */
header('Cache-Control: no-cache,must-revalidate');
header('Pragma: no-cache');
header("Expires: 0"); 
header('P3P:CP=CAO PSA OUR');//解决跨域不传cookie

ini_set("var","value");//设置本次访问的全局变量var=value
define("c","value");//定义一个常量
include_once '../xx.php'; //执行一次x.php,相对路径基于调用着,不安全  dirname(__FILE__)获得本文件路径

error_reporting(E_ALL & ~E_NOTICE);//设置php的全局错误处理 位运算
set_error_handler("error_handler");//设置用户处理error方法
error_get_last();//获得最后返回的error

register_shutdown_function("callback");//访问结束按注册顺序调用回调,回调中调用exit直接返回,不再触发后续回调




//--------------util------------------

if (isset($_GET["varX"])){//查询本次访问get传递变量varX是否存在且非空
    //do something
};


json_decode($ret);//把ret json序列化

empty($x);//x是否为空

preg_match("/0[1-9]/",$var);//正则匹配

class_exists("pdo");// 检查类是否已定义