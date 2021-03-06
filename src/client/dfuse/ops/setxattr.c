/**
 * (C) Copyright 2019-2021 Intel Corporation.
 *
 * SPDX-License-Identifier: BSD-2-Clause-Patent
 */

#include "dfuse_common.h"
#include "dfuse.h"

#include "daos_uns.h"

void
dfuse_cb_setxattr(fuse_req_t req, struct dfuse_inode_entry *inode,
		  const char *name, const char *value, size_t size,
		  int flags)
{
	int rc;

	DFUSE_TRA_DEBUG(inode, "Attribute '%s'", name);

	if (strcmp(name, DUNS_XATTR_NAME) == 0) {
		struct duns_attr_t	dattr = {};

		rc = duns_parse_attr((char *)value, size, &dattr);
		if (rc)
			D_GOTO(err, rc);
	}

	rc = dfs_setxattr(inode->ie_dfs->dfs_ns, inode->ie_obj, name, value,
			  size, flags);
	if (rc == 0) {
		DFUSE_REPLY_ZERO(inode, req);
		return;
	}
err:
	DFUSE_REPLY_ERR_RAW(inode, req, rc);
}
