import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders} from '@angular/common/http';

@Injectable()
export class HttpClientService {
  /**
   * HTTP クライアントを実行する際のヘッダオプション
   * 認証トークンを使用するために httpOptions としてオブジェクトを用意
   */
  public httpOptions: any = {
    // ヘッダ情報
    headers: new HttpHeaders({
      'Content-Type': 'application/json',
      // 'Access-Controle-Allow-Origin': 'http://localhost:8080'
    }),
    body: null,
  };

  /**
   * API 実行時に指定する URL
   * バックエンドは Go 言語で実装し、8080番ポートで待ち受けている
   * TODO: ハードコードされているのは今後修正
   */
  public host = 'http://localhost:8080';

  /**
   * HttpClientService のインスタンスを生成するコンストラクタ
   */
  constructor(private http: HttpClient) {
    this.setAuthorization('my-auth-token');
  }

  /**
   * Authorization に認証トークンを設定する
   * トークンを動的に設定するためにメソッド化した
   */
  public setAuthorization(token: string = null): void {
    if (!token) {
      return;
    }
    const bearerToken = `Bearer {$token}`;
    this.httpOptions.headers = this.httpOptions.headers.set('Authorizarion', bearerToken);
  }
}
