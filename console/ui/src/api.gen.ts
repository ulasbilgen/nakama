// tslint:disable
/* Code generated by openapi-gen/main.go. DO NOT EDIT. */

const BASE_PATH = "http://127.0.0.1:80";

export interface ConfigurationParameters {
  basePath?: string;
  username?: string;
  password?: string;
  bearerToken?: string;
  timeoutMs?: number;
}
/** A warning for a configuration field. */
export interface ConfigWarning {
  // The config field this warning is for in a JSON pointer format.
  field?: string;
  // Warning message text.
  message?: string;
}
/** The status of a Nakama node. */
export interface StatusListStatus {
  // Average input bandwidth usage.
  avg_input_kbs?: number;
  // Average response latency in milliseconds.
  avg_latency_ms?: number;
  // Average output bandwidth usage.
  avg_output_kbs?: number;
  // Average number of requests per second.
  avg_rate_sec?: number;
  // Current number of running goroutines.
  goroutine_count?: number;
  // Health score.
  health?: number;
  // Current number of active authoritative matches.
  match_count?: number;
  // Node name.
  name?: string;
  // Currently registered live presences.
  presence_count?: number;
  // Currently connected sessions.
  session_count?: number;
}
/** A single group-role pair. */
export interface UserGroupListUserGroup {
  // Group.
  group?: ApiGroup;
  // The user's relationship to the group.
  state?: number;
}
/** A user with additional account details. Always the current user. */
export interface ApiAccount {
  // The custom id in the user's account.
  custom_id?: string;
  // The devices which belong to the user's account.
  devices?: Array<ApiAccountDevice>;
  // The email address of the user.
  email?: string;
  // The user object.
  user?: ApiUser;
  // The UNIX time when the user's email was verified.
  verify_time?: string;
  // The user's wallet data.
  wallet?: string;
}
/** Send a device to the server. Used with authenticate/link/unlink and user. */
export interface ApiAccountDevice {
  // A device identifier. Should be obtained by a platform-specific device API.
  id?: string;
}
/** A message sent on a channel. */
export interface ApiChannelMessage {
  // The channel this message belongs to.
  channel_id?: string;
  // The code representing a message type or category.
  code?: number;
  // The content payload.
  content?: string;
  // The UNIX time when the message was created.
  create_time?: string;
  // The unique ID of this message.
  message_id?: string;
  // True if the message was persisted to the channel's history, false otherwise.
  persistent?: boolean;
  // Message sender, usually a user ID.
  sender_id?: string;
  // The UNIX time when the message was last updated.
  update_time?: string;
  // The username of the message sender, if any.
  username?: string;
}
/** A friend of a user. */
export interface ApiFriend {
  // The friend status.
  state?: number;
  // The user object.
  user?: ApiUser;
}
/** A collection of zero or more friends of the user. */
export interface ApiFriends {
  // The Friend objects.
  friends?: Array<ApiFriend>;
}
/** A group in the server. */
export interface ApiGroup {
  // A URL for an avatar image.
  avatar_url?: string;
  // The UNIX time when the group was created.
  create_time?: string;
  // The id of the user who created the group.
  creator_id?: string;
  // A description for the group.
  description?: string;
  // The current count of all members in the group.
  edge_count?: number;
  // The id of a group.
  id?: string;
  // The language expected to be a tag which follows the BCP-47 spec.
  lang_tag?: string;
  // The maximum number of members allowed.
  max_count?: number;
  // Additional information stored as a JSON object.
  metadata?: string;
  // The unique name of the group.
  name?: string;
  // Anyone can join open groups, otherwise only admins can accept members.
  open?: boolean;
  // The UNIX time when the group was last updated.
  update_time?: string;
}
/** Represents a complete leaderboard record with all scores and associated metadata. */
export interface ApiLeaderboardRecord {
  // The UNIX time when the leaderboard record was created.
  create_time?: string;
  // The UNIX time when the leaderboard record expires.
  expiry_time?: string;
  // The ID of the leaderboard this score belongs to.
  leaderboard_id?: string;
  // The maximum number of score updates allowed by the owner.
  max_num_score?: number;
  // Metadata.
  metadata?: string;
  // The number of submissions to this score record.
  num_score?: number;
  // The ID of the score owner, usually a user or group.
  owner_id?: string;
  // The rank of this record.
  rank?: string;
  // The score value.
  score?: string;
  // An optional subscore value.
  subscore?: string;
  // The UNIX time when the leaderboard record was updated.
  update_time?: string;
  // The username of the score owner, if the owner is a user.
  username?: string;
}
/** A notification in the server. */
export interface ApiNotification {
  // Category code for this notification.
  code?: number;
  // Content of the notification in JSON.
  content?: string;
  // The UNIX time when the notification was created.
  create_time?: string;
  // ID of the Notification.
  id?: string;
  // True if this notification was persisted to the database.
  persistent?: boolean;
  // ID of the sender, if a user. Otherwise 'null'.
  sender_id?: string;
  // Subject of the notification.
  subject?: string;
}
/** An object within the storage engine. */
export interface ApiStorageObject {
  // The collection which stores the object.
  collection?: string;
  // The UNIX time when the object was created.
  create_time?: string;
  // The key of the object within the collection.
  key?: string;
  // The read access permissions for the object.
  permission_read?: number;
  // The write access permissions for the object.
  permission_write?: number;
  // The UNIX time when the object was last updated.
  update_time?: string;
  // The user owner of the object.
  user_id?: string;
  // The value of the object.
  value?: string;
  // The version hash of the object.
  version?: string;
}
/** A storage acknowledgement. */
export interface ApiStorageObjectAck {
  // The collection which stores the object.
  collection?: string;
  // The key of the object within the collection.
  key?: string;
  // The owner of the object.
  user_id?: string;
  // The version hash of the object.
  version?: string;
}
/** A user in the server. */
export interface ApiUser {
  // A URL for an avatar image.
  avatar_url?: string;
  // The UNIX time when the user was created.
  create_time?: string;
  // The display name of the user.
  display_name?: string;
  // Number of related edges to this user.
  edge_count?: number;
  // The Facebook id in the user's account.
  facebook_id?: string;
  // The Apple Game Center in of the user's account.
  gamecenter_id?: string;
  // The Google id in the user's account.
  google_id?: string;
  // The id of the user's account.
  id?: string;
  // The language expected to be a tag which follows the BCP-47 spec.
  lang_tag?: string;
  // The location set by the user.
  location?: string;
  // Additional information stored as a JSON object.
  metadata?: string;
  // Indicates whether the user is currently online.
  online?: boolean;
  // The Steam id in the user's account.
  steam_id?: string;
  // The timezone set by the user.
  timezone?: string;
  // The UNIX time when the user was last updated.
  update_time?: string;
  // The username of the user's account.
  username?: string;
}
/** A list of groups belonging to a user, along with the user's role in each group. */
export interface ApiUserGroupList {
  // Group-role pairs for a user.
  user_groups?: Array<UserGroupListUserGroup>;
}
/** An export of all information stored for a user account. */
export interface ConsoleAccountExport {
  // The user's account details.
  account?: ApiAccount;
  // The user's friends.
  friends?: Array<ApiFriend>;
  // The user's groups.
  groups?: Array<ApiGroup>;
  // The user's leaderboard records.
  leaderboard_records?: Array<ApiLeaderboardRecord>;
  // The user's chat messages.
  messages?: Array<ApiChannelMessage>;
  // The user's notifications.
  notifications?: Array<ApiNotification>;
  // The user's storage.
  objects?: Array<ApiStorageObject>;
  // The user's wallet ledger items.
  wallet_ledgers?: Array<ConsoleWalletLedger>;
}
/** Authenticate a console user with username and password. */
export interface ConsoleAuthenticateRequest {
  // The password of the user.
  password?: string;
  // The username of the user.
  username?: string;
}
/** The current server configuration and any associated warnings. */
export interface ConsoleConfig {
  // JSON-encoded active server configuration.
  config?: string;
  // Any warnings about the current config.
  warnings?: Array<ConfigWarning>;
}
/** A console user session. */
export interface ConsoleConsoleSession {
  // A session token (JWT) for the console user.
  token?: string;
}
/** List of nodes and their stats. */
export interface ConsoleStatusList {
  // List of nodes and their stats.
  nodes?: Array<StatusListStatus>;
}
/** List of storage objects. */
export interface ConsoleStorageList {
  // List of storage objects matching list/filter operation.
  objects?: Array<ApiStorageObject>;
  // Approximate total number of storage objects.
  total_count?: number;
}
/** A list of users. */
export interface ConsoleUserList {
  // Approximate total number of users.
  total_count?: number;
  // A list of users.
  users?: Array<ApiUser>;
}
/** An individual update to a user's wallet. */
export interface ConsoleWalletLedger {
  // The changeset.
  changeset?: string;
  // The UNIX time when the wallet ledger item was created.
  create_time?: string;
  // The identifier of this wallet change.
  id?: string;
  // Any associated metadata.
  metadata?: string;
  // The UNIX time when the wallet ledger item was updated.
  update_time?: string;
  // The user ID this wallet ledger item belongs to.
  user_id?: string;
}
/** List of wallet ledger items for a particular user. */
export interface ConsoleWalletLedgerList {
  // A list of wallet ledger items.
  items?: Array<ConsoleWalletLedger>;
}

export const NakamaApi = (configuration: ConfigurationParameters = {
  basePath: BASE_PATH,
  bearerToken: "",
  password: "",
  username: "",
  timeoutMs: 5000,
}) => {
  return {
    /** Perform the underlying Fetch operation and return Promise object **/
    doFetch(urlPath: string, method: string, queryParams: any, body?: any, options?: any): Promise<any> {
      const urlQuery = "?" + Object.keys(queryParams)
        .map(k => {
          if (queryParams[k] instanceof Array) {
            return queryParams[k].reduce((prev: any, curr: any) => {
              return prev + encodeURIComponent(k) + "=" + encodeURIComponent(curr) + "&";
            }, "");
          } else {
            if (queryParams[k] != null) {
              return encodeURIComponent(k) + "=" + encodeURIComponent(queryParams[k]) + "&";
            }
          }
        })
        .join("");

      const fetchOptions = {...{ method: method /*, keepalive: true */ }, ...options};
      const headers = {
        "Accept": "application/json",
        "Content-Type": "application/json",
      } as any;

      if (configuration.bearerToken) {
        headers["Authorization"] = "Bearer " + configuration.bearerToken;
      } else if (configuration.username) {
        headers["Authorization"] = "Basic " + btoa(configuration.username + ":" + configuration.password);
      }

      fetchOptions.headers = {...headers, ...options.headers};
      fetchOptions.body = body;

      return Promise.race([
        fetch(configuration.basePath + urlPath + urlQuery, fetchOptions).then((response) => {
          if (response.status >= 200 && response.status < 300) {
            return response.json();
          } else {
            throw response;
          }
        }),
        new Promise((_, reject) =>
          setTimeout(reject, configuration.timeoutMs, "Request timed out.")
        ),
      ]);
    },
    /** Delete all information stored for a user account. */
    deleteAccount(id: string, recordDeletion?: boolean, options: any = {}): Promise<any> {
      if (id === null || id === undefined) {
        throw new Error("'id' is a required parameter but is null or undefined.");
      }
      const urlPath = "/v2/console/account/{id}"
         .replace("{id}", encodeURIComponent(String(id)));

      const queryParams = {
        record_deletion: recordDeletion,
      } as any;

      let _body = null;

      return this.doFetch(urlPath, "DELETE", queryParams, _body, options)
    },
    /** Get detailed account information for a single user. */
    getAccount(id: string, options: any = {}): Promise<ApiAccount> {
      if (id === null || id === undefined) {
        throw new Error("'id' is a required parameter but is null or undefined.");
      }
      const urlPath = "/v2/console/account/{id}"
         .replace("{id}", encodeURIComponent(String(id)));

      const queryParams = {
      } as any;

      let _body = null;

      return this.doFetch(urlPath, "GET", queryParams, _body, options)
    },
    /** Update one or more fields on a user account. */
    updateAccount(id: string, options: any = {}): Promise<any> {
      if (id === null || id === undefined) {
        throw new Error("'id' is a required parameter but is null or undefined.");
      }
      const urlPath = "/v2/console/account/{id}"
         .replace("{id}", encodeURIComponent(String(id)));

      const queryParams = {
      } as any;

      let _body = null;

      return this.doFetch(urlPath, "POST", queryParams, _body, options)
    },
    /** Ban a user. */
    banUser(id: string, options: any = {}): Promise<any> {
      if (id === null || id === undefined) {
        throw new Error("'id' is a required parameter but is null or undefined.");
      }
      const urlPath = "/v2/console/account/{id}/ban"
         .replace("{id}", encodeURIComponent(String(id)));

      const queryParams = {
      } as any;

      let _body = null;

      return this.doFetch(urlPath, "POST", queryParams, _body, options)
    },
    /** Export all information stored about a user account. */
    exportAccount(id: string, options: any = {}): Promise<ConsoleAccountExport> {
      if (id === null || id === undefined) {
        throw new Error("'id' is a required parameter but is null or undefined.");
      }
      const urlPath = "/v2/console/account/{id}/export"
         .replace("{id}", encodeURIComponent(String(id)));

      const queryParams = {
      } as any;

      let _body = null;

      return this.doFetch(urlPath, "GET", queryParams, _body, options)
    },
    /** Get a user's list of friend relationships. */
    getFriends(id: string, options: any = {}): Promise<ApiFriends> {
      if (id === null || id === undefined) {
        throw new Error("'id' is a required parameter but is null or undefined.");
      }
      const urlPath = "/v2/console/account/{id}/friend"
         .replace("{id}", encodeURIComponent(String(id)));

      const queryParams = {
      } as any;

      let _body = null;

      return this.doFetch(urlPath, "GET", queryParams, _body, options)
    },
    /** Delete the friend relationship between two users. */
    deleteFriend(id: string, friendId: string, options: any = {}): Promise<any> {
      if (id === null || id === undefined) {
        throw new Error("'id' is a required parameter but is null or undefined.");
      }
      if (friendId === null || friendId === undefined) {
        throw new Error("'friendId' is a required parameter but is null or undefined.");
      }
      const urlPath = "/v2/console/account/{id}/friend/{friend_id}"
         .replace("{id}", encodeURIComponent(String(id)))
         .replace("{friend_id}", encodeURIComponent(String(friendId)));

      const queryParams = {
      } as any;

      let _body = null;

      return this.doFetch(urlPath, "DELETE", queryParams, _body, options)
    },
    /** Get a list of groups the user is a member of. */
    getGroups(id: string, options: any = {}): Promise<ApiUserGroupList> {
      if (id === null || id === undefined) {
        throw new Error("'id' is a required parameter but is null or undefined.");
      }
      const urlPath = "/v2/console/account/{id}/group"
         .replace("{id}", encodeURIComponent(String(id)));

      const queryParams = {
      } as any;

      let _body = null;

      return this.doFetch(urlPath, "GET", queryParams, _body, options)
    },
    /** Remove a user from a group. */
    deleteGroupUser(id: string, groupId: string, options: any = {}): Promise<any> {
      if (id === null || id === undefined) {
        throw new Error("'id' is a required parameter but is null or undefined.");
      }
      if (groupId === null || groupId === undefined) {
        throw new Error("'groupId' is a required parameter but is null or undefined.");
      }
      const urlPath = "/v2/console/account/{id}/group/{group_id}"
         .replace("{id}", encodeURIComponent(String(id)))
         .replace("{group_id}", encodeURIComponent(String(groupId)));

      const queryParams = {
      } as any;

      let _body = null;

      return this.doFetch(urlPath, "DELETE", queryParams, _body, options)
    },
    /** Unban a user. */
    unbanUser(id: string, options: any = {}): Promise<any> {
      if (id === null || id === undefined) {
        throw new Error("'id' is a required parameter but is null or undefined.");
      }
      const urlPath = "/v2/console/account/{id}/unban"
         .replace("{id}", encodeURIComponent(String(id)));

      const queryParams = {
      } as any;

      let _body = null;

      return this.doFetch(urlPath, "POST", queryParams, _body, options)
    },
    /** Unlink the custom ID from a user account. */
    unlinkCustom(id: string, options: any = {}): Promise<any> {
      if (id === null || id === undefined) {
        throw new Error("'id' is a required parameter but is null or undefined.");
      }
      const urlPath = "/v2/console/account/{id}/unlink/custom"
         .replace("{id}", encodeURIComponent(String(id)));

      const queryParams = {
      } as any;

      let _body = null;

      return this.doFetch(urlPath, "POST", queryParams, _body, options)
    },
    /** Unlink the device ID from a user account. */
    unlinkDevice(id: string, options: any = {}): Promise<any> {
      if (id === null || id === undefined) {
        throw new Error("'id' is a required parameter but is null or undefined.");
      }
      const urlPath = "/v2/console/account/{id}/unlink/device"
         .replace("{id}", encodeURIComponent(String(id)));

      const queryParams = {
      } as any;

      let _body = null;

      return this.doFetch(urlPath, "POST", queryParams, _body, options)
    },
    /** Unlink the email from a user account. */
    unlinkEmail(id: string, options: any = {}): Promise<any> {
      if (id === null || id === undefined) {
        throw new Error("'id' is a required parameter but is null or undefined.");
      }
      const urlPath = "/v2/console/account/{id}/unlink/email"
         .replace("{id}", encodeURIComponent(String(id)));

      const queryParams = {
      } as any;

      let _body = null;

      return this.doFetch(urlPath, "POST", queryParams, _body, options)
    },
    /** Unlink the Facebook ID from a user account. */
    unlinkFacebook(id: string, options: any = {}): Promise<any> {
      if (id === null || id === undefined) {
        throw new Error("'id' is a required parameter but is null or undefined.");
      }
      const urlPath = "/v2/console/account/{id}/unlink/facebook"
         .replace("{id}", encodeURIComponent(String(id)));

      const queryParams = {
      } as any;

      let _body = null;

      return this.doFetch(urlPath, "POST", queryParams, _body, options)
    },
    /** Unlink the Game Center ID from a user account. */
    unlinkGameCenter(id: string, options: any = {}): Promise<any> {
      if (id === null || id === undefined) {
        throw new Error("'id' is a required parameter but is null or undefined.");
      }
      const urlPath = "/v2/console/account/{id}/unlink/gamecenter"
         .replace("{id}", encodeURIComponent(String(id)));

      const queryParams = {
      } as any;

      let _body = null;

      return this.doFetch(urlPath, "POST", queryParams, _body, options)
    },
    /** Unlink the Google ID from a user account. */
    unlinkGoogle(id: string, options: any = {}): Promise<any> {
      if (id === null || id === undefined) {
        throw new Error("'id' is a required parameter but is null or undefined.");
      }
      const urlPath = "/v2/console/account/{id}/unlink/google"
         .replace("{id}", encodeURIComponent(String(id)));

      const queryParams = {
      } as any;

      let _body = null;

      return this.doFetch(urlPath, "POST", queryParams, _body, options)
    },
    /** Unlink the Steam ID from a user account. */
    unlinkSteam(id: string, options: any = {}): Promise<any> {
      if (id === null || id === undefined) {
        throw new Error("'id' is a required parameter but is null or undefined.");
      }
      const urlPath = "/v2/console/account/{id}/unlink/steam"
         .replace("{id}", encodeURIComponent(String(id)));

      const queryParams = {
      } as any;

      let _body = null;

      return this.doFetch(urlPath, "POST", queryParams, _body, options)
    },
    /** Get a list of the user's wallet transactions. */
    getWalletLedger(id: string, options: any = {}): Promise<ConsoleWalletLedgerList> {
      if (id === null || id === undefined) {
        throw new Error("'id' is a required parameter but is null or undefined.");
      }
      const urlPath = "/v2/console/account/{id}/wallet"
         .replace("{id}", encodeURIComponent(String(id)));

      const queryParams = {
      } as any;

      let _body = null;

      return this.doFetch(urlPath, "GET", queryParams, _body, options)
    },
    /** Delete a wallet ledger item. */
    deleteWalletLedger(id: string, walletId: string, options: any = {}): Promise<any> {
      if (id === null || id === undefined) {
        throw new Error("'id' is a required parameter but is null or undefined.");
      }
      if (walletId === null || walletId === undefined) {
        throw new Error("'walletId' is a required parameter but is null or undefined.");
      }
      const urlPath = "/v2/console/account/{id}/wallet/{wallet_id}"
         .replace("{id}", encodeURIComponent(String(id)))
         .replace("{wallet_id}", encodeURIComponent(String(walletId)));

      const queryParams = {
      } as any;

      let _body = null;

      return this.doFetch(urlPath, "DELETE", queryParams, _body, options)
    },
    /** Authenticate a console user with username and password. */
    authenticate(body: ConsoleAuthenticateRequest, options: any = {}): Promise<ConsoleConsoleSession> {
      if (body === null || body === undefined) {
        throw new Error("'body' is a required parameter but is null or undefined.");
      }
      const urlPath = "/v2/console/authenticate";

      const queryParams = {
      } as any;

      let _body = null;
      _body = JSON.stringify(body || {});

      return this.doFetch(urlPath, "POST", queryParams, _body, options)
    },
    /** Get server config and configuration warnings. */
    getConfig(options: any = {}): Promise<ConsoleConfig> {
      const urlPath = "/v2/console/config";

      const queryParams = {
      } as any;

      let _body = null;

      return this.doFetch(urlPath, "GET", queryParams, _body, options)
    },
    /** Get current status data for all nodes. */
    getStatus(options: any = {}): Promise<ConsoleStatusList> {
      const urlPath = "/v2/console/status";

      const queryParams = {
      } as any;

      let _body = null;

      return this.doFetch(urlPath, "GET", queryParams, _body, options)
    },
    /** Delete all storage data. */
    deleteStorage(options: any = {}): Promise<any> {
      const urlPath = "/v2/console/storage";

      const queryParams = {
      } as any;

      let _body = null;

      return this.doFetch(urlPath, "DELETE", queryParams, _body, options)
    },
    /** List (and optionally filter) storage data. */
    listStorage(userId?: string, options: any = {}): Promise<ConsoleStorageList> {
      const urlPath = "/v2/console/storage";

      const queryParams = {
        user_id: userId,
      } as any;

      let _body = null;

      return this.doFetch(urlPath, "GET", queryParams, _body, options)
    },
    /** Delete a storage object. */
    deleteStorageObject(collection: string, key: string, userId: string, version?: string, options: any = {}): Promise<any> {
      if (collection === null || collection === undefined) {
        throw new Error("'collection' is a required parameter but is null or undefined.");
      }
      if (key === null || key === undefined) {
        throw new Error("'key' is a required parameter but is null or undefined.");
      }
      if (userId === null || userId === undefined) {
        throw new Error("'userId' is a required parameter but is null or undefined.");
      }
      const urlPath = "/v2/console/storage/{collection}/{key}/{user_id}"
         .replace("{collection}", encodeURIComponent(String(collection)))
         .replace("{key}", encodeURIComponent(String(key)))
         .replace("{user_id}", encodeURIComponent(String(userId)));

      const queryParams = {
        version: version,
      } as any;

      let _body = null;

      return this.doFetch(urlPath, "DELETE", queryParams, _body, options)
    },
    /** Get a storage object. */
    getStorage(collection: string, key: string, userId: string, options: any = {}): Promise<ApiStorageObject> {
      if (collection === null || collection === undefined) {
        throw new Error("'collection' is a required parameter but is null or undefined.");
      }
      if (key === null || key === undefined) {
        throw new Error("'key' is a required parameter but is null or undefined.");
      }
      if (userId === null || userId === undefined) {
        throw new Error("'userId' is a required parameter but is null or undefined.");
      }
      const urlPath = "/v2/console/storage/{collection}/{key}/{user_id}"
         .replace("{collection}", encodeURIComponent(String(collection)))
         .replace("{key}", encodeURIComponent(String(key)))
         .replace("{user_id}", encodeURIComponent(String(userId)));

      const queryParams = {
      } as any;

      let _body = null;

      return this.doFetch(urlPath, "GET", queryParams, _body, options)
    },
    /** Write a new storage object or replace an existing one. */
    writeStorageObject(collection: string, key: string, userId: string, options: any = {}): Promise<ApiStorageObjectAck> {
      if (collection === null || collection === undefined) {
        throw new Error("'collection' is a required parameter but is null or undefined.");
      }
      if (key === null || key === undefined) {
        throw new Error("'key' is a required parameter but is null or undefined.");
      }
      if (userId === null || userId === undefined) {
        throw new Error("'userId' is a required parameter but is null or undefined.");
      }
      const urlPath = "/v2/console/storage/{collection}/{key}/{user_id}"
         .replace("{collection}", encodeURIComponent(String(collection)))
         .replace("{key}", encodeURIComponent(String(key)))
         .replace("{user_id}", encodeURIComponent(String(userId)));

      const queryParams = {
      } as any;

      let _body = null;

      return this.doFetch(urlPath, "POST", queryParams, _body, options)
    },
    /** Delete a storage object. */
    deleteStorageObject2(collection: string, key: string, userId: string, version: string, options: any = {}): Promise<any> {
      if (collection === null || collection === undefined) {
        throw new Error("'collection' is a required parameter but is null or undefined.");
      }
      if (key === null || key === undefined) {
        throw new Error("'key' is a required parameter but is null or undefined.");
      }
      if (userId === null || userId === undefined) {
        throw new Error("'userId' is a required parameter but is null or undefined.");
      }
      if (version === null || version === undefined) {
        throw new Error("'version' is a required parameter but is null or undefined.");
      }
      const urlPath = "/v2/console/storage/{collection}/{key}/{user_id}/{version}"
         .replace("{collection}", encodeURIComponent(String(collection)))
         .replace("{key}", encodeURIComponent(String(key)))
         .replace("{user_id}", encodeURIComponent(String(userId)))
         .replace("{version}", encodeURIComponent(String(version)));

      const queryParams = {
      } as any;

      let _body = null;

      return this.doFetch(urlPath, "DELETE", queryParams, _body, options)
    },
    /** Delete (non-recorded) all user accounts. */
    deleteUsers(options: any = {}): Promise<any> {
      const urlPath = "/v2/console/user";

      const queryParams = {
      } as any;

      let _body = null;

      return this.doFetch(urlPath, "DELETE", queryParams, _body, options)
    },
    /** List (and optionally filter) users. */
    listUsers(filter?: string, banned?: boolean, tombstones?: boolean, options: any = {}): Promise<ConsoleUserList> {
      const urlPath = "/v2/console/user";

      const queryParams = {
        filter: filter,
        banned: banned,
        tombstones: tombstones,
      } as any;

      let _body = null;

      return this.doFetch(urlPath, "GET", queryParams, _body, options)
    },
  };
};
