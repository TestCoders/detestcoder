package com.training;

public class Main {

    public static void main(String[] args) {
	    var customerRepository = new InMemoryCustomerRepository();
	    var customerService = new CustomerService(customerRepository);

	    var customer = new Customer(1, "Mauro", "Palsgraaf");

	    System.out.println(customerService.getCustomer(1));

	    customerService.save(customer);

	    System.out.println(customerService.getCustomer(1));
    }
}
