// 引入axios
import axios from 'axios'

// 创建axios实例，添加api路径
const http = axios.create({
    baseURL:'http://localhost:7777/admin/api'
})

// 将文件变量导出，导出到main.js
export default http
