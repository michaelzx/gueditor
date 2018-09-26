package common

const (
    BIGGER_THAN_UPLOAD_MAX_FILESIZE string = "文件大小超出 upload_max_filesize 限制"
    BIGGER_THAN_MAX_FILE_SIZE       string = "文件大小超出 MAX_FILE_SIZE 限制"
    FILE_NOT_COMPLETE               string = "文件未被完整上传"
    NO_FILE_UPLOAD                  string = "没有文件被上传"
    UPLOAD_FILE_IS_EMPTY            string = "上传文件为空"
    ERROR_TMP_FILE                  string = "临时文件错误"
    ERROR_TMP_FILE_NOT_FOUND        string = "找不到临时文件"
    ERROR_SIZE_EXCEED               string = "文件大小超出网站限制"
    ERROR_TYPE_NOT_ALLOWED          string = "文件类型不允许"
    ERROR_CREATE_DIR                string = "目录创建失败"
    ERROR_DIR_NOT_WRITEABLE         string = "目录没有写权限"
    ERROR_FILE_MOVE                 string = "文件保存时出错"
    ERROR_FILE_NOT_FOUND            string = "找不到上传文件"
    ERROR_WRITE_CONTENT             string = "写入文件内容错误"
    ERROR_UNKNOWN                   string = "未知错误"
    ERROR_DEAD_LINK                 string = "链接不可用"
    ERROR_HTTP_LINK                 string = "链接不是http链接"
    ERROR_HTTP_CONTENTTYPE          string = "链接contentType不正确"
    INVALID_URL                     string = "非法 URL"
    INVALID_IP                      string = "非法 IP"
    ERROR_BASE64_DATA               string = "base64图片解码错误"
    ERROR_FILE_STATE                string = "文件系统错误"
    ERRPR_READ_REMOTE_DATA          string = "读取远程链接出错"
)