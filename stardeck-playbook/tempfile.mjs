import * as fs from 'node:fs';

import { temporaryFileTask } from 'tempy';
import yaml from 'yaml';

export function temporaryYamlFileTask(data, callback) {
  return temporaryFileTask((tempPath) => {
    fs.writeFileSync(tempPath, yaml.stringify(data));
    callback(tempPath);
  });
}
