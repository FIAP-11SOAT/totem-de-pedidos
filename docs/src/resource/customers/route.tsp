import "@typespec/http";
import "@typespec/openapi";

import "./models.tsp";

using TypeSpec.Http;

namespace main.customers;

@route("/customers")
@tag("Customers")
interface Customers {
    
    @get()
    @summary("Get all customers")
    @doc("Get all customers")
    getCustomers(): Customer[];

    @post()
    @summary("Create a new customer")
    @doc("Create a new customer")
    createCustomer(customer: Customer): Customer;
    
}