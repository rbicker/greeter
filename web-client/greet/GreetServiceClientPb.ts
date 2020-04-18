/**
 * @fileoverview gRPC-Web generated client stub for greet
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


import * as grpcWeb from 'grpc-web';

import {
  GreetEveryoneRequest,
  GreetEveryoneResponse,
  GreetManyTimesRequest,
  GreetManyTimesResponse,
  GreetRequest,
  GreetResponse,
  GreetWithDeadlineRequest,
  GreetWithDeadlineResponse,
  LongGreetRequest,
  LongGreetResponse} from './greet_pb';

export class GreetServiceClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: string; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: string; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodInfoGreet = new grpcWeb.AbstractClientBase.MethodInfo(
    GreetResponse,
    (request: GreetRequest) => {
      return request.serializeBinary();
    },
    GreetResponse.deserializeBinary
  );

  greet(
    request: GreetRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: GreetResponse) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/greet.GreetService/Greet',
      request,
      metadata || {},
      this.methodInfoGreet,
      callback);
  }

  methodInfoGreetManyTimes = new grpcWeb.AbstractClientBase.MethodInfo(
    GreetManyTimesResponse,
    (request: GreetManyTimesRequest) => {
      return request.serializeBinary();
    },
    GreetManyTimesResponse.deserializeBinary
  );

  greetManyTimes(
    request: GreetManyTimesRequest,
    metadata?: grpcWeb.Metadata) {
    return this.client_.serverStreaming(
      this.hostname_ +
        '/greet.GreetService/GreetManyTimes',
      request,
      metadata || {},
      this.methodInfoGreetManyTimes);
  }

  methodInfoGreetWithDeadline = new grpcWeb.AbstractClientBase.MethodInfo(
    GreetWithDeadlineResponse,
    (request: GreetWithDeadlineRequest) => {
      return request.serializeBinary();
    },
    GreetWithDeadlineResponse.deserializeBinary
  );

  greetWithDeadline(
    request: GreetWithDeadlineRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: GreetWithDeadlineResponse) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/greet.GreetService/GreetWithDeadline',
      request,
      metadata || {},
      this.methodInfoGreetWithDeadline,
      callback);
  }

}

