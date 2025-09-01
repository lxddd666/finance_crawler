import { http, jumpExport } from '@/utils/http/axios';

// 获取股票代码列表
export function List(params) {
  return http.request({
    url: '/financeCode/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除股票代码
export function Delete(params) {
  return http.request({
    url: '/financeCode/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑股票代码
export function Edit(params) {
  return http.request({
    url: '/financeCode/edit',
    method: 'POST',
    params,
  });
}

// 获取股票代码指定详情
export function View(params) {
  return http.request({
    url: '/financeCode/view',
    method: 'GET',
    params,
  });
}

// 导出股票代码
export function Export(params) {
  jumpExport('/financeCode/export', params);
}