# 执行python脚本,生成新的tag
tag:
	python3 tag.py


# 要操作的tag标签
tag := $(firstword $(MAKECMDGOALS))
# 删除tag,推送到远程
del: del_local
	git push origin :refs/tags/$(tag)

# 先删除本地tag
.PHONY: del_local
del_local:
	git tag -d $(tag)
	@echo "delete tag $(tag)"