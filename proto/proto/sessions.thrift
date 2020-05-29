namespace go sessions

service Waiter {
  ResponseAddPddSessions AddPddSessions (1: RequestAddPddSessions requestAddPddSessions);
  ResponseUpdatePddSessions UpdatePddSessions (1: RequestUpdatePddSessions requestUpdatePddSessions);
  ResponseDeletePddSessions DeletePddSessions (1: RequestDeletePddSessions requestDeletePddSessions);
  ResponseGetPddSessionsInfo GetPddSessionsInfo (1: RequestGetPddSessionsInfo requestGetPddSessionsInfo);
  ResponseGetPddSessionsList GetPddSessionsList (1: RequestGetPddSessionsList requestGetPddSessionsList);
}


struct PddSessions {
    1: optional i32 id;
    2: optional i32 taokeID;
    3: optional string screenName;
    4: optional string openId;
    5: optional string token;
    6: optional i64 expiredAt;
    7: optional string refreshToken;
    8: optional i64 refreshExpiredAt;
    9: optional i32 isDefault;
    10: optional i64 createdAt;
    11: optional i64 updatedAt;
}


struct ResponseGetPddSessionsInfo {
  1: optional PddSessions pddSessions;
}

struct RequestGetPddSessionsInfo {
  1: optional i32 id;
}

struct ResponseGetPddSessionsList {
  1: optional list<PddSessions> pddSessions;
  2: optional i32 total;
}

struct RequestGetPddSessionsList {
  1: optional i32 page = 1;
  2: optional i32 per_page = 10;
}

struct RequestAddPddSessions {
    1: optional PddSessions pddSessions;
}

struct ResponseAddPddSessions {
    1: optional PddSessions pddSessions;
}

struct RequestUpdatePddSessions {
    1: optional PddSessions pddSessions;
}

struct ResponseUpdatePddSessions {
    1: optional PddSessions pddSessions;
}

struct RequestDeletePddSessions {
  1: optional i32 id;
}

struct ResponseDeletePddSessions {
}