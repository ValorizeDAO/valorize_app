export default interface Ethereum {
  send(arg0: string): Promise<Array<string>>;
  enable(): Promise<any>;
  request(args: RequestArguments): Promise<unknown>;
}

interface RequestArguments {
  readonly method: string;
  readonly params?: readonly unknown[] | object;
}

interface ProviderRpcError extends Error {
  code: number;
  data?: unknown;
}

interface ProviderMessage {
  readonly type: string;
  readonly data: unknown;
}