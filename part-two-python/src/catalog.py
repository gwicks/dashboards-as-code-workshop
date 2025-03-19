import typing


class Service:
    name: str
    description: str
    has_http: bool
    has_grpc: bool
    repository_url: str

    def __init__(self, name: str = "", description: str = "", has_http: bool = False, has_grpc: bool = False, repository_url: str = ""):
        self.name = name
        self.description = description
        self.has_http = has_http
        self.has_grpc = has_grpc
        self.repository_url = repository_url

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "name" in data:
            args["name"] = data["name"]
        if "description" in data:
            args["description"] = data["description"]
        if "has_http" in data:
            args["has_http"] = data["has_http"]   
        if "has_grpc" in data:
            args["has_grpc"] = data["has_grpc"]   
        if "repository_url" in data:
            args["repository_url"] = data["github"]        

        return cls(**args)
