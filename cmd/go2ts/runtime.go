package main

const runtime = `
export function identityBoolean(x: boolean): boolean {
  return x;
}

export function identityNumber(x: number): number {
  return x;
}

export function identityString(x: string): string {
  return x;
}

export function fmapArray<T, U>(fn: (t: T) => U): (a: Array<T> | null) => Array<U> {
  return function(x) {
    if (x === null) {
      return [];
    }
    return x.map(fn);
  };
}

export function fmapObject<T, U>(fn: (t: T) => U): (m: { [k: string]: T } | null) => { [k: string]: U } {
  return function(x) {
    var retval = {};
    if (x === null) {
      return retval;
    }
    for (var key in x) {
      retval[key] = fn(x[key]);
    }
    return retval;
  };
}

export function bytesUnmarshal(x: string | null): string {
  if (!x) {
    return "";
  }
  return window.atob(x);
}
`
