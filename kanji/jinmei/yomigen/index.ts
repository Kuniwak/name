import {DB} from "https://deno.land/x/sqlite@v3.9.1/mod.ts";
import {strokeMap} from '../stroke.ts';
import {KanjiYomiEntry} from '../../types/yomi.ts';

const db = new DB("./moji.db");

const result: KanjiYomiEntry[] = [];
for (const kanji of strokeMap.keys()) {
  const res = db.query("SELECT 読み FROM mji INNER JOIN main.mji_reading mr ON mji.MJ文字図形名 = mr.MJ文字図形名 WHERE 実装したUCS = ? OR 対応するUCS = ?;", [kanji, kanji])
  const yomi = res.map((row) => row[0] as string).filter((yomi) => yomi.match(/^[ァ-ヴー]+$/));
  result.push({ kanji, yomi });
}

db.close();

console.log(JSON.stringify(result, null, "  "));