import { createNamespace, Namespace } from 'cls-hooked';
import { Tracing } from './constants';

class Context {
  private _ns: Namespace;
  private _context: any;
  constructor(ns: string) {
    this._ns = createNamespace(ns);
    this._context = this._ns.createContext();
  }

  private _bindContext = (cb: any) => {
    return this._ns.bind(cb, this._context);
  };

  public set = (key: string, value: string) => {
    const setContext = this._bindContext(() => this._ns.set(key, value));
    setContext();
  };
  public get = (key: string) => {
    const getContext = this._bindContext(() => this._ns.get(key));
    return getContext();
  };
}

export default new Context(Tracing.TRACER_SESSION);
