namespace go Exvs

struct College {
    1: required string name(go.tag = 'json:"name"'),
    2: string address(go.tag = 'json:"address"'),
}

struct Student {
    1: required i32 id,
    2: required string name,
    3: required College college,
    4: optional list<string> email,
}

struct RegisterResp {
    1: bool success,
    2: string message,
}

struct QueryReq {
    1: required i32 id (api.query="id"),
}

struct UpdateReq {
    1: string version
}

struct UpdateResp {
    1: bool Success
    2: string idl
}

//----------------------service-------------------
service StudentService {
    RegisterResp Register(1: Student info) (api.post="/add-student-info")
    Student Query(1: QueryReq req) (api.get="/query")
    UpdateResp Update(1: UpdateReq req) (api.post="/update")
}