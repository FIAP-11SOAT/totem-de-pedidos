import "@typespec/http";
import "@typespec/openapi";

import "./models.tsp";

using TypeSpec.Http;
using TypeSpec.OpenAPI;

namespace main.customers;

@route("/customers")
@tag("Customers")
interface Customers {
    
    @summary("Get all customers")
    @doc("Get all customers")
    @get getCustomers(): Customer[];

    @summary("Get a customer by ID")
    @doc("Get a customer by ID")
    @get getCustomer(@path id: string): Customer;

    @summary("Create a new customer")
    @doc("Create a new customer")
    @post createCustomer(customer: Customer): Customer;

    @summary("Delete a customer")
    @doc("Delete a customer")
    @delete deleteCustomer(@path id: string): void;

}