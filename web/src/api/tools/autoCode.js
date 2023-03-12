import { request } from '@/utils/request.js'


export const preview = (data) => {
  return request({
    url: '/autoCode/preview',
    method: 'post',
    data
  })
}

export const createTemp = (data) => {
  return request({
    url: '/autoCode/createTemp',
    method: 'post',
    data,
    responseType: 'blob'
  })
}

// @Tags SysApi
// @Summary 获取当前所有数据库
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /autoCode/getDatabase [get]
export const getDB = (params) => {
  return request({
    url: '/autoCode/getDB',
    method: 'get',
    params
  })
}

// @Tags SysApi
// @Summary 获取当前数据库所有表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /autoCode/getTables [get]
export const getTable = (params) => {
  return request({
    url: '/autoCode/getTables',
    method: 'get',
    params
  })
}

// @Tags SysApi
// @Summary 获取当前数据库所有表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /autoCode/getColumn [get]
export const getColumn = (params) => {
  return request({
    url: '/autoCode/getColumn',
    method: 'get',
    params
  })
}

export const getSysHistory = (data) => {
  return request({
    url: '/autoCode/getSysHistory',
    method: 'post',
    data
  })
}

export const rollback = (data) => {
  return request({
    url: '/autoCode/rollback',
    method: 'post',
    data
  })
}

export const getMeta = (data) => {
  return request({
    url: '/autoCode/getMeta',
    method: 'post',
    data
  })
}

export const delSysHistory = (data) => {
  return request({
    url: '/autoCode/delSysHistory',
    method: 'post',
    data
  })
}

export const createPackageApi = (data) => {
  return request({
    url: '/autoCode/createPackage',
    method: 'post',
    data
  })
}

export const getPackageApi = () => {
  return request({
    url: '/autoCode/getPackage',
    method: 'post'
  })
}

export const deletePackageApi = (data) => {
  return request({
    url: '/autoCode/delPackage',
    method: 'post',
    data
  })
}

export const createPlugApi = (data) => {
  return request({
    url: '/autoCode/createPlug',
    method: 'post',
    data
  })
}

export const installPlug = (data) => {
  return request({
    url: '/autoCode/installPlug',
    method: 'post',
    data
  })
}
