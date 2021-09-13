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

$x= "";

ini_set("var","value");//设置本次访问的全局变量var=value
define("c","value");//定义一个常量,仅限本次访问
$_SESSION["USER"]="user1";//定义会话可访问的数组内容,再次连接也可取值
include_once '../xx.php'; //执行一次x.php,相对路径基于调用着,不安全  dirname(__FILE__)获得本文件路径
error_reporting(E_ALL & ~E_NOTICE);//设置php的全局错误处理 位运算
set_error_handler("error_handler");//设置用户处理error方法
error_get_last();//获得最后返回的error
register_shutdown_function("callback");//访问结束按注册顺序调用回调,回调中调用exit直接返回,不再触发后续回调
function_exists("funcname");//function_exists — 如果给定的函数已经被定义就返回 true
die("32131");//die() 函数输出一条消息，并退出当前脚本
ob_start();// 打开输出控制缓冲,可以定义callback  https://www.php.net/manual/zh/function.ob-start.php
//范围解析操作符 （::） 范围解析操作符（也可称作 Paamayim Nekudotayim）或者更简单地说是一对冒号，可以用于访问静态成员，类常量，还可以用于覆盖类中的属性和方法。
spl_autoload_register();//spl_autoload_register — 注册给定的函数作为 __autoload 的实现
ucwords("1");//ucwords — 将字符串中每个单词的首字母转换为大写
strpos("",".");//strpos — 查找字符串首次出现的位置
is_numeric($x);//is_numeric — 检测变量是否为数字或数字字符串
php_strip_whitespace("filename");//php_strip_whitespace — 返回删除注释和空格后的PHP源码
trim("");//trim — 去除字符串首尾处的空白字符（或者其他字符）
substr("str","startindex","len");//返回字符串的子串
array_merge(array(),array());//array_merge — 合并一个或多个数组
is_array($x);//var 是否为数组
get_defined_constants();//get_defined_constants — 返回所有常量的关联数组，键是常量名，值是常量值 这包含 define() 函数所创建的，也包含了所有扩展所创建的。
if (isset($_GET["varX"])){//查询本次访问get传递变量varX是否存在且非空
    //do something
};


json_decode($ret);//把ret json序列化

empty($x);//x是否为空

preg_match("/0[1-9]/",$var);//正则匹配

class_exists("pdo");// 检查类是否已定义