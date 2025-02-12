function colorCode(code) {
  return `\u001b[${code}m`;
}

const RESET = colorCode(0);

export const LEVELS = {
  debug: { severity: 40, verbosity: 3, color: 35 },
  verbose: { tag: 'verb', severity: 30, verbosiry: 2, color: 36 },
  info: { severity: 20, verbosity: 1, color: 32 },
  warn: { severity: 10, verbosity: 0, color: 33 },
  error: { severity: 0, verbosity: 0, color: 31 },
};

export const LEVEL_ALIASES = {
  warning: 'warn',
};

const MAX_LEN = Math.max(
  ...Object.entries(LEVELS).map(([name, lvl]) => (lvl.tag || name).length),
);

for (let [name, lvl] of Object.entries(LEVELS)) {
  lvl.padding = MAX_LEN - name.length;
}

let logger = {
  _log: console.error,
  _severity: 10,

  log(level, message, ...optionalParams) {
    if (LEVELS[level].severity <= this._severity) {
      const tag = LEVELS[level].tag || level;
      const color = colorCode(LEVELS[level].color);
      let msg = `${color}${tag}${RESET}`;
      msg = msg + ':' + ' '.repeat(LEVELS[level].padding);

      if (typeof message == 'string') {
        this._log(`${msg} ${message}`, ...optionalParams);
      } else {
        this._log(msg, message, ...optionalParams);
      }
    }
  },

  setLevel(lvl) {
    this._severity = LEVELS[lvl].severity;
  },
};

for (let lvl of Object.keys(LEVELS)) {
  logger[lvl] = function (message, ...optionalParams) {
    this.log(lvl, message, ...optionalParams);
  };
}

for (let [alias, to] of Object.entries(LEVEL_ALIASES)) {
  logger[alias] = logger[to];
}

export { logger };
