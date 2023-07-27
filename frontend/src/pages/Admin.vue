<template>
  <div class="w-full h-full flex flex-grow">
    <div class="flex flex-col w-52 bg-gray-800 ">
      <router-link class="h-10 flex items-center justify-center" active-class="bg-gray-600" :to="'/admin/contests'">コンテスト</router-link>
      <router-link class="h-10 flex items-center justify-center"  active-class="bg-gray-600" :to="'/admin/users'">ユーザ</router-link>
    </div>
    <div class="flex flex-col flex-grow">
      <div class="h-full m-5">
        <router-view></router-view>
      </div>
    </div>
  </div>
</template> 
