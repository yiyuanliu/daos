/*
 * (C) Copyright 2018-2020 Intel Corporation.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * GOVERNMENT LICENSE RIGHTS-OPEN SOURCE SOFTWARE
 * The Government's rights to use, modify, reproduce, release, perform, display,
 * or disclose this software are subject to the terms of the Apache License as
 * provided in Contract No. B609815.
 * Any reproduction of computer software, computer software documentation, or
 * portions thereof marked with this legend must also reproduce the markings.
 */

package io.daos.dfs.uns;

public interface DaosAceOrBuilder extends
      // @@protoc_insertion_point(interface_extends:uns.DaosAce)
      com.google.protobuf.MessageOrBuilder {

  /**
   * <code>uint32 access_types = 1;</code>
   *
   * @return The accessTypes.
   */
  int getAccessTypes();

  /**
   * <code>uint32 principal_type = 2;</code>
   *
   * @return The principalType.
   */
  int getPrincipalType();

  /**
   * <code>uint32 principal_len = 3;</code>
   *
   * @return The principalLen.
   */
  int getPrincipalLen();

  /**
   * <code>uint32 access_flags = 4;</code>
   *
   * @return The accessFlags.
   */
  int getAccessFlags();

  /**
   * <code>uint32 reserved = 5;</code>
   *
   * @return The reserved.
   */
  int getReserved();

  /**
   * <code>uint32 allow_perms = 6;</code>
   *
   * @return The allowPerms.
   */
  int getAllowPerms();

  /**
   * <code>uint32 audit_perms = 7;</code>
   *
   * @return The auditPerms.
   */
  int getAuditPerms();

  /**
   * <code>uint32 alarm_perms = 8;</code>
   *
   * @return The alarmPerms.
   */
  int getAlarmPerms();

  /**
   * <code>string principal = 9;</code>
   *
   * @return The principal.
   */
  java.lang.String getPrincipal();

  /**
   * <code>string principal = 9;</code>
   *
   * @return The bytes for principal.
   */
  com.google.protobuf.ByteString
      getPrincipalBytes();
}