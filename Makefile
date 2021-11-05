PATCH_DIR	= patch

patch:
	cd patch && sh patch.sh patch.json

.PHONY:	patch