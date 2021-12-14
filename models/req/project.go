/*
 * @Date: 2021-12-14 10:17:25
 * @LastEditTime: 2021-12-14 12:31:54
 * @Author: ceaqw
 */
package req

type CreateProjectForm struct {
	Id          uint64 `json:"id"`
	Platform    string `json:"platform"`
	Project     string `json:"project"`
	ScriptId    uint64 `json:"scriptId"`
	Ip          string `json:"ip"`
	CriticalKpi string `json:"criticalKpi"`
	Server      uint8  `json:"server"`
	Comment     string `json:"comment"`
	BindTables  string `json:"bindTables"`
}
