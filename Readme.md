# 微服务示例
将thrift文件编译为go文件

thrift -r --gen go echo.thrift

基本语法

基本类型

bool: 布尔值 对应Java中的boolean 

byte: 有符号字节 对应Java中的byte 

i16: 16位有符号整型 对应Java中的short 

i32: 32位有符号整型 对应Java中的int 

i64: 64位有符号整型 对应Java中的long 

double: 64位浮点型 对应Java中的double 

string: 字符串 对应Java中的String 

binary: Blob 类型 对应Java中的byte[]

容器类型

集合中的元素可以是除了service之外的任何类型，包括exception。

list<T>: 一系列由T类型的数据组成的有序列表，元素可以重复

set<T>: 一系列由T类型的数据组成的无序集合，元素不可重复

map<K, V>: 一个字典结构，key为K类型，value为V类型，相当于Java中的HashMap<K,V>

结构体

就像C语言一样，thrift也支持struct类型，目的就是将一些数据聚合在一起，方便传输管理。struct的定义形式如下：

struct NPC

{

	1:i32 id;
	
	2:string name; 
	
}

 枚举
 
枚举的定义形式和Java的Enum定义差不多，例如：

enum Action {

    Idle,
    
      Attack,
      
    Run 
    
}

异常

thrift支持自定义exception，规则和struct一样，如下：

exception RequestException {

    1: i32 code;
    
    2: string reason;
    
}

服务

thrift定义服务相当于Java中创建Interface一样，创建的service经过代码生成命令之后就会生成客户端和服务端的框架代码。定义形式如下：

service HelloWordService {

     // service中定义的函数，相当于Java interface中定义的函数
     
     string doAction(1: string name, 2: i32 age);
     
 }
 
类型定义

 thrift支持类似C++一样的typedef定义，比如：

typedef i32 Integer

typedef i64 Long

注意：末尾没有逗号或者分号！

常量

thrift也支持常量定义，使用const关键字，例如：

const i32 MAX_RETRIES_TIME = 10;

const string MY_WEBSITE = "http://qifuguang.me";

末尾的分号是可选的，可有可无，并且支持16进制赋值


命名空间

thrift的命名空间相当于Java中的package的意思，主要目的是组织代码。thrift使用关键字namespace定义命名空间，例如：

namespace java com.game.lll.thrift

提示：格式是namespace 语言(Java)  路径(com.game.lll.thrift)， 注意末尾不能有分号。

文件包含

thrift也支持文件包含，相当于C/C++中的include，Java中的import,C#中的using。使用关键字include定义，例 如：

include "global.thrift"

注释

 thrift注释方式支持shell风格的注释，支持C/C++风格的注释，即#和//开头的语句都单当做注释，/**/包裹的语句也是注释。

可选与必选

thrift提供两个关键字required，optional，分别用于表示对应的字段时必填的还是可选的。例如：

struct People {

    1: required string name;
    
    2: optional i32 age;
    
}

表示name是必填的，age是可选的。

thrift编译

步骤一：创建一个文件，代码如下：

namespace java com.game.lll.thrift

struct Request {

    1: string username;   
       
    2: string password;  
             
}

exception RequestException {

    1: required i32 code;
    
    2: optional string reason;
    
}

// 服务名

service LoginService {

    string doAction(1: Request request) throws (1:RequestException qe); // 可能抛出异常。
    
}

步骤二：在终端输入命令thrift -gen java login.thrift后会在当前目录下生成gen-java文件夹，该文件夹下会按照namespace定义的路径名一次一层层生成文件夹,到gen-java/com/game/lll/thrift/目录下可以看到生成的3个.java类。