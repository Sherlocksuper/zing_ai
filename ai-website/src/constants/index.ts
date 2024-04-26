import {Setting} from "../utils/settingUtils.ts";

export const KimiToken = "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ1c2VyLWNlbnRlciIsImV4cCI6MTcxMzAwNjk0MCwiaWF0IjoxNzEzMDA2MDQwLCJqdGkiOiJjb2Q2Ym0xa3FxNHR0cm5odm1vMCIsInR5cCI6ImFjY2VzcyIsInN1YiI6ImNvNGlsbDZjcDdmMmhwMThrODBnIiwic3BhY2VfaWQiOiJjbzRpbGw2Y3A3ZjJocDE4azgwMCIsImFic3RyYWN0X3VzZXJfaWQiOiJjbzRpbGw2Y3A3ZjJocDE4azd2ZyJ9.hs65uAx4hszu7DO_pwSsi8N1YxsdJzKI9ZC-YAWzrjQYDgFJmfO2lbqQOD4Aajnbmw61DHaaJ51Ogm3cwvvMRw"

export const defaultSetting: Setting = {
    model: "kimi",
    use_search: true,
    stream: false
}

interface settingItem {
    title: string,
    func: () => void
}

export const settingList: settingItem[] = [
    {
        title: "暗夜模式",
        func: () => {
        }
    },
    {
        title: "是否联网",
        func: () => {
            console.log("Use Search")
        }
    },
]