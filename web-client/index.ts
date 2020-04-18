import * as grpcWeb from 'grpc-web';
import {GreetServiceClient} from './greet/GreetServiceClientPb'
import {Greeting, GreetRequest, GreetResponse} from "./greet/greet_pb";

const client = new GreetServiceClient('https://greeter-web.asdf.ooo');
const request = new GreetRequest();
const greeting = new Greeting();
greeting.setFirstName('Raphael');
greeting.setLastName('Bicker')
request.setGreeting(greeting);

const call = client.greet(request, {'custom-header-1': 'value1'},
    (err: grpcWeb.Error, response: GreetResponse) => {
        console.log(response.getResult());
        console.log(err);
    });
call.on('status', (status: grpcWeb.Status) => {
    console.log(status);
});