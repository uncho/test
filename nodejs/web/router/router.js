// 要为路由提供请求的 URL 和其他需要的 GET 及 POST 参数，随后路由需要根据这些数据来执行相应的代码

function route(pathname){
    console.log('the request\'s path is : ' + pathname)
}

exports.route = route;