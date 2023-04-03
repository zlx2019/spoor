import subprocess

tag = ''
# tag前缀
tag_prefix = "v"
# tag分隔符
tag_sep = "."
# 初始化版本
tag_first_version = "v1.0.0"


# 执行终端命令函数
# commands: 要执行的命令数组
# timeout: 命令超时时间,默认为不限制
# encoding: 标准输入与输出的编码格式
def execute(*commands, timeout=None, encoding="utf-8"):
    # 命令执行结果
    # 只有return_code为0 才表示正常执行成功
    process_result = subprocess.run(commands, capture_output=True, text=True, encoding=encoding, timeout=timeout)
    if process_result.returncode == 0:
        return True, process_result.stdout.splitlines()
    else:
        # 执行失败,返回False和错误信息
        return False, process_result.stderr


ok, result = execute("git", "tag")
if ok:
    last_tag = ""
    if len(result) > 0:
        # 获取最后一个tag
        last_tag = result[-1]
        # 去除前缀`v` 和`.`符号,获取版本号
        last_tag_num = last_tag.replace(tag_prefix, "").replace(tag_sep, "")
        # 生成下一个tag版本号
        new_tag_num = int(last_tag_num) + 1
        num_list = list(str(new_tag_num))
        new_tag = tag_sep.join(num_list)
        tag = "%s%s" % (tag_prefix, new_tag)
    else:
        # 使用第一个tag版本
        tag = tag_first_version
    print("%s --> %s" % (last_tag, tag))
    # 创建tag
    ok, result = execute("git", "tag", tag)
    if ok:
        print("Crate Tag %s Success" % tag)
        # 推送tag
        # ok, result = execute("git", "tag", "origin", tag)
        if ok:
            print("Push Tag %s Success" % tag)
        # 推送tag失败
        else:
            print("Push Tag %s Fail: %s" % (tag, result))
    # 创建tag失败
    else:
        print("Create Tag %s Fail: %s" % (tag, result))
else:
    print(result)