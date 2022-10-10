import { defineConfig } from 'vitepress'

export default defineConfig({
    title: "kpm",
    description: '快速的，节省磁盘空间的包管理工具', 
    base: '/', 
    lang: 'zh-Hans-CN', 
    outDir: '../dist',
    lastUpdated: true,
    appearance: true,
    markdown: {
        lineNumbers: true
    },
    
    // 主题配置
    themeConfig: {

        //   头部导航
        nav: [
            { text: '文档', link: '/guide/' },
            { text: 'GitHub', link: 'https://github.com/orangebees/kpm' },

        ],
        //   侧边导航
        sidebar: [
            {
                text: '介绍',
                items: [
                    { text: '项目背景', link: '/guide/' },
                    { text: '安装', link: '/guide/installation' },
                    // { text: '功能比较', link: '/guide/' },
                ]
            },
            {
                text: 'CLI命令',
                items: [
                    { text: 'kpm add <pkg>', link: '/guide/cli/add' },
                    { text: 'kpm del <pkg>', link: '/guide/cli/del' },
                    { text: 'kpm download', link: '/guide/cli/download' },
                    { text: 'kpm graph', link: '/guide/cli/graph' },
                    { text: 'kpm init <pkg>', link: '/guide/cli/init' },
                    { text: 'kpm publish', link: '/guide/cli/publish' },
                    { text: 'kpm search', link: '/guide/cli/search' },
                    { text: 'kpm tidy', link: '/guide/cli/tidy' },
                    { text: 'kpm store add <pkg>', link: '/guide/cli/store/add' },
                    { text: 'kpm store addfile <path>', link: '/guide/cli/store/addfile' },
                ]
            },
            {
                text: '配置',
                items: [
                    { text: 'kpm.json', link: '/guide/configuring/kpm.json' },
                    { text: 'hash.csv', link: '/guide/configuring/hash.csv' },
                    { text: 'info.json', link: '/guide/configuring/info.json' },
                    { text: 'dir', link: '/guide/configuring/dir' },
                    { text: 'env', link: '/guide/configuring/env' },
                ]
            },

        ],
        footer: {
            message: 'Released under the  Apache-2.0 license License.',
            copyright: 'Copyright © 2022-present orangebees'
          },
          docFooter: {
            prev: '上一章',
            next: '下一章'
          }
    }
})
