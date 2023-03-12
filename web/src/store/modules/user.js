import {defineStore} from 'pinia'
import loginApi from '@/api/login'
import tool from '@/utils/tool'
import router from '@/router'
import webRouter from '@/router/webRouter'
import {homePage} from '@/router/homePageRoutes'
import {useAppStore} from '@/store'

const useUserStore = defineStore('user', {

    state: () => ({
        codes: undefined,
        roles: undefined,
        routers: undefined,
        user: undefined,
        menus: undefined,
    }),

    getters: {
        setUserInfo(state) {
            return {...state};
        },
    },

    actions: {

        setToken(token) {
            tool.local.set(import.meta.env.VITE_APP_TOKEN_PREFIX, token)
        },
        setUser(user) {
            tool.local.set("xuserid", user.ID)
        },
        getToken() {
            return tool.local.get(import.meta.env.VITE_APP_TOKEN_PREFIX)
        },

        clearToken() {
            tool.local.remove(import.meta.env.VITE_APP_TOKEN_PREFIX)
        },

        setInfo(data) {
            this.$patch(data)
        },

        resetUserInfo() {
            this.$reset();
        },

        setMenu(data) {
            const routers = flatAsyncRoutes(filterAsyncRouter(data))

            console.log("fuck",routers)
            routers.map(item => router.addRoute('layout', item))
        },

        requestUserInfo() {
            return new Promise((resolve, reject) => {
                loginApi.getInfo().then(response => {

                   // console.log(!response , !response.data,!response || !response.data)
                    if (!response || !response.data) {
                        console(3)
                        router.push({name: 'login'})
                        return
                    }
                    let data = response.data
                    data.codes = [
                        "*"
                    ]
                    data.roles = ["superAdmin"]
                    this.setInfo(data)

                    homePage.children = webRouter[0].children

                    this.setMenu(this.routers)

                    this.routers = removeButtonMenu(this.routers)

                    this.routers.unshift(homePage)
                    this.setApp()
                    resolve(data)
                }).catch(_ => {
                    this.clearToken()
                    router.push({name: 'login'})
                })
            }).catch(error => {
                this.clearToken()
                router.push({name: 'login'})
                reject(error)
            })

        },

        login(form) {
            return loginApi.login(form).then(r => {
                if (r.code == 0) {
                    this.setToken(r.data.token)
                    this.setUser(r.data.user)
                    return true
                } else {
                    return false
                }
            }).catch(e => {
                console.error(e)
                return false
            })
        },

        async logout() {
            await loginApi.logout()
            tool.local.remove('tags')
            this.clearToken()
            this.resetUserInfo()
        },

        async setApp() {
            const appStore = useAppStore()

            console.log("this.user")
            console.log(this.user)
            const setting = this.user.backend_setting
            if (setting) {
                appStore.toggleMode(setting.mode)
                appStore.toggleMenu(setting.menuCollapse)
                appStore.toggleTag(setting.tag)
                appStore.changeMenuWidth(setting.menuWidth)
                appStore.changeLayout(setting.layout)
                appStore.useSkin(setting.skin)
                appStore.changeColor(setting.color)
            }
        }
    }

})

//路由扁平化
const flatAsyncRoutes = (routes, breadcrumb = []) => {
    let res = []
    routes.forEach(route => {
        const tmp = {...route}
        if (tmp.children) {
            let childrenBreadcrumb = [...breadcrumb]
            childrenBreadcrumb.push(route)
            let tmpRoute = {...route}
            tmpRoute.meta.breadcrumb = childrenBreadcrumb
            delete tmpRoute.children
            res.push(tmpRoute)
            let childrenRoutes = flatAsyncRoutes(tmp.children, childrenBreadcrumb)
            childrenRoutes.map(item => {
                res.push(item)
            })
        } else {
            let tmpBreadcrumb = [...breadcrumb]
            tmpBreadcrumb.push(tmp)
            tmp.meta.breadcrumb = tmpBreadcrumb
            res.push(tmp)
        }
    })

    console.log("==]",res)
    return res
}

const views = import.meta.glob('../../views/**/**.vue')
const empty = import.meta.glob('../../layout/empty.vue')

// 菜单转换路由
const filterAsyncRouter = (routerMap) => {
    const accessedRouters = []
    routerMap.forEach(item => {
        item.path=item.route
        item.meta.title=item.name
        if (item.meta.type !== 'B') {

            if (item.meta.type === 'I') {
                item.meta.url = item.path
                item.path = `/maIframe/${item.name}`
            }
            item.hidden=item.meta.is_hidden
            const route = {
                path: item.path,
                name: item.name,
                hidden: item.hidden == 1,
                meta: item.meta,
                children: item.children ? filterAsyncRouter(item.children) : null,
                component: views[`../../views/${item.component}.vue`]
            }
           // console.log(route)
            accessedRouters.push(route)
        }
    })
    return accessedRouters
}

// 去除按钮菜单
const removeButtonMenu = (routers) => {
    let handlerAfterRouters = []
    routers.forEach(item => {
        item.meta.hidden=item.meta.is_hidden
        if (item.meta.type !== 'B' && !item.meta.hidden) {
            let route = item
            if (item.children && item.children.length > 0) {
                route.children = removeButtonMenu(item.children)
            }
            handlerAfterRouters.push(route)
        }
    })
    return handlerAfterRouters
}
export default useUserStore
