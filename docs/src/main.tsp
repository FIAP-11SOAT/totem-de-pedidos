import "@typespec/http";
import "@typespec/openapi";

import "./constants.tsp";
import "./resource/customers/route.tsp";

using TypeSpec.OpenAPI;


@service()
@info(Info)
@doc(DescMain)
namespace main;


