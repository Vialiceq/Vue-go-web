<template>
    <el-table :data="tableData">
            <el-table-column prop="date" label="日期" width="140">
            </el-table-column>
            <el-table-column prop="name" label="姓名" width="120">
            </el-table-column>
            <el-table-column prop="address" label="地址">
            </el-table-column>
    </el-table>
</template>

<script>
  export default {
    data() {
      const item = {
        date: '2022-05-20',
        name: 'alice',
        address: 'university of oxford'
      };
      return {
        tableData: Array(25).fill(item)
      }
    }
  };
</script>