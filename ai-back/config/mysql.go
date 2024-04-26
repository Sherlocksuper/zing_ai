package config

import (
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io/ioutil"
	"os/exec"
	"strconv"
	"time"
)

type Mysql struct {
	Host       string `json:"host" yaml:"host"`         // 服务器IP地址
	Port       int    `json:"port" yaml:"port"`         // 服务器端口号
	UserName   string `json:"username" yaml:"username"` // 数据库用户名
	Password   string `json:"password" yaml:"password"` // 数据库用户密码
	Database   string `json:"database" yaml:"database"` // 数据库名
	Prefix     string `json:"prefix" yaml:"prefix"`     // 数据表前缀
	Charset    string `json:"charset" yaml:"charset"`   // 字符集
	BackupPath string `json:"backup_path" yaml:"backup_path"`
}

// SetHost 设置Host
func (m *Mysql) SetHost(host string) *Mysql {
	m.Host = host
	return m
}

// SetPort 设置Port
func (m *Mysql) SetPort(port int) *Mysql {
	m.Port = port
	return m
}

// SetUserName 设置UserName
func (m *Mysql) SetUserName(userName string) *Mysql {
	m.UserName = userName
	return m
}

// SetPassword 设置Password
func (m *Mysql) SetPassword(password string) *Mysql {
	m.Password = password
	return m
}

// SetDatabase 设置Database
func (m *Mysql) SetDatabase(database string) *Mysql {
	m.Database = database
	return m

}

// SetPrefix 设置Prefix
func (m *Mysql) SetPrefix(prefix string) *Mysql {
	m.Prefix = prefix
	return m
}

// SetCharset 设置Charset
func (m *Mysql) SetCharset(charset string) *Mysql {
	m.Charset = charset
	return m
}

// SetBackupPath 设置BackupPath
func (m *Mysql) SetBackupPath(backupPath string) *Mysql {
	m.BackupPath = backupPath
	return m
}

func (m *Mysql) GetDb() (*gorm.DB, error) {
	DSN := m.Dsn()
	var Db, _ = gorm.Open(mysql.Open(DSN), &gorm.Config{})
	return Db, nil
}

func (m *Mysql) Dsn() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		m.UserName,
		m.Password,
		m.Host,
		m.Port,
		m.Database,
		m.Charset,
	)
}

// CheckError 检查是否有错误
func (m *Mysql) CheckError() error {
	if m.Host == "" {
		return errors.New("host不能为空")
	}

	if m.Port == 0 {
		log.Error().Msg("port不能为空")
		return errors.New("port不能为空")
	}

	if m.UserName == "" {
		log.Error().Msg("userName不能为空")
		return errors.New("userName不能为空")
	}

	if m.Password == "" {
		log.Error().Msg("password不能为空")
		return errors.New("password不能为空")
	}

	if m.Database == "" {
		log.Error().Msg("database不能为空")
		return errors.New("database不能为空")
	}

	if m.BackupPath == "" {
		log.Error().Msg("BackupPath不能为空")
		return errors.New("BackupPath不能为空")
	}
	return nil
}

// BackupDb  备份
func (m *Mysql) BackupDb() error {
	err := m.CheckError()
	if err != nil {
		return err
	}
	err, _ = BackupMySqlDb(m.Host, strconv.Itoa(m.Port), m.UserName, m.Password, m.Database, "", m.BackupPath)
	if err != nil {
		log.Error().Msg(err.Error())
		return err
	}
	return nil
}

// BackupTable 备份表
func (m *Mysql) BackupTable(tableName string) error {
	err := m.CheckError()
	if err != nil {
		log.Error().Msg(err.Error())
		return err
	}
	err, _ = BackupMySqlDb(m.Host, string(rune(m.Port)), m.UserName, m.Password, m.Database, tableName, m.BackupPath)
	return nil
}

// BackupMySqlDb 备份 MySQL 数据库
func BackupMySqlDb(host, port, user, password, databaseName, tableName, sqlPath string) (error, string) {
	var cmd *exec.Cmd

	if tableName == "" {
		//打印出执行的命令
		log.Info().Msg("mysqldump --opt -h" + host + " -P" + port + " -u" + user + " -p" + password + " " + databaseName)
		cmd = exec.Command("mysqldump", "--opt", "-h"+host, "-P"+port, "-u"+user, "-p"+password, databaseName)
	} else {
		cmd = exec.Command("mysqldump", "--opt", "-h"+host, "-P"+port, "-u"+user, "-p"+password, databaseName, tableName)
	}

	stdout, err := cmd.StdoutPipe()

	if err != nil {
		log.Error().Msg(err.Error())
		return err, ""
	}

	if err := cmd.Start(); err != nil {
		log.Error().Msg(err.Error())
		return err, ""
	}

	bytes, err := ioutil.ReadAll(stdout)

	if err != nil {
		log.Error().Msg(err.Error())
		return err, ""
	}

	now := time.Now().Format("20060102150405")
	var backupPath string

	if tableName == "" {
		backupPath = sqlPath + databaseName + "_" + now + ".sql"
	} else {
		backupPath = sqlPath + databaseName + "_" + tableName + "_" + now + ".sql"
	}

	err = ioutil.WriteFile(backupPath, bytes, 0644)
	if err != nil {
		log.Error().Msg(err.Error())
		return err, ""
	}

	return nil, backupPath
}
