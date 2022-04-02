<template>
    <div>
        <h1>分类列表</h1>
        <el-table :data="items">
            <el-table-column prop="_id" label="ID" width="220">
            </el-table-column>
            <el-table-column prop="parent.name" label="上级分类">
            </el-table-column>
            <el-table-column prop="name" label="分类名称">
            </el-table-column>
            <el-table-column
            fixed="right"
            label="操作"
            width="100">
                <template slot-scope="scope">
                    <el-button type="text" size="small" @click="$router.push('/categories/edit/' + scope.row._id)">编辑</el-button>
                    <el-button @click="remove(scope.row)" type="text" size="small">删除</el-button>
                </template>
            </el-table-column>
        </el-table>
    </div>
</template>
<script>
export default {
    data() {
        return {
            items: []
        }
    },
    methods: {
        async fetch(){
            const res = await this.$http.get('rest/categories')
            this.items = res.data
        },
        remove(row){
            this.$confirm('是否确定要删除分类"' + row.name + '"?', '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning'
            }).then(async () => {
                // 要想使用await，函数必须使用async
                // await异步执行，待调用接口获取数据完成后再将值传给res，进行下一步操作
                const res = await this.$http.delete('rest/categories/' + row._id)
                this.$message({
                    type: 'success',
                    message: '删除成功!'
                });
                if(res.status == 200){
                    // 接口调用成功后，刷新页面
                    this.fetch()
                }
            }).catch(() => {
                this.$message({
                    type: 'info',
                    message: '已取消删除'
                });          
            });
        }
    },
    created() {
        this.fetch()
    }
}
</script>