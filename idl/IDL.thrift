namespace go Exvs

struct UpdateReq {
    1: string version
}

struct UpdateResp {
    1: bool Success
    2: string idl
}

//----------------------service-------------------
service StudentService {
    UpdateResp Update(1: UpdateReq req) (api.post="/update")
}