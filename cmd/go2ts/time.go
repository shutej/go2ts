package main

const time_Time = `
let partialTime = (
    /^(\d{4})-(\d{2})-(\d{2})T(\d{2}):(\d{2}):(\d{2})(?:\.(\d{3}))?/
);

let timeNumOffset = (
    /^Z|([+-]\d{2}):(\d{2})$/
);

export type T = Date;

export function empty(): T {
  let tmp = new Date(Date.UTC(0, 0, 1, 0, 0, 0, 0));  // 1900AD
  tmp.setUTCFullYear(1);  // 1AD
  return tmp;
}

export type MarshalT = string;

export function unmarshal(s: MarshalT): T {
  let m = partialTime.exec(s);
  if (!m) {
    throw "unable to parse input as RFC3339";
  }

  let year     = +m[1];
  let month    = +m[2];
  let day      = +m[3];
  let hour     = +m[4];
  let minute   = +m[5];
  let second   = +m[6];
  let msec     = +m[7] || 0;

  s = s.substr(m[0].length);
  m = timeNumOffset.exec(s);
  if (!m) {
    throw "unable to parse input as RFC3339";
  }

  let tzHour = +m[1] || 0;
  let tzMin  = +m[2] || 0;

  // This is the timezone offset of the client clock.
  let tzOffset = tzHour * 60 + tzMin;

  let retval = new Date(Date.UTC(0, month - 1, day, hour, minute - tzOffset, second, msec));
  // Adjust for full year; date offset will be incorporated, above.
  retval.setUTCFullYear(retval.getUTCFullYear() - 1900 + year);
  return retval;
}

export function marshal(d: T): MarshalT {
  function pad(pad, n) {
    let str = "" + n;
    return pad.substring(0, pad.length - str.length) + str;
  }

  return pad("0000", d.getUTCFullYear()) + "-"
    +  pad("00", d.getUTCMonth() + 1) + "-"
    +  pad("00", d.getUTCDate()) + "T"
    +  pad("00", d.getUTCHours()) + ":"
    +  pad("00", d.getUTCMinutes()) + ":"
    +  pad("00", d.getUTCSeconds()) + "Z";
}
`
