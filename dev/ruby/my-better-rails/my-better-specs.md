# Request Spec
```rb
RSpec.describe 'Hoges Pages', type: :request do
  let(:user)  { create(:user) }
  let!(:hoge) { create(:hoge, user: user) }

  describe 'GET /hoges' do
    context 'as an authenticated user' do # 認証ユーザとして
      before { sign_in user }

      context 'with authorized access' do # 許可されたアクセス
        before { get hoges_path }

        it 'responds with 200' do # 200 で応答する
          expect(response).to have_http_status '200'
        end

        context 'when responds JSON is correct' do # 応答した JSON が正しい
          let(:json_response) { JSON.parse(response.body, symbolize_names: true) }

          it 'has correct keys' do # 正しいキーを持つか
            expect(json_response[0].keys).to include :id, :name
          end
          it 'has correct values' do # 正しい値を持つか
            expect(json_response[0].values).to include hoge.id, hoge.name
          end
        end
      end
    end

    context 'as an unauthenticated user' do # 認証していないユーザとして
      context 'with unauthorized access' do # 許可されていないアクセス
        it 'responds with 302' do # 302 で応答する
        end
        it 'redirects to root_path' do # root_path にリダイレクトする
        end
      end
    end
  end
end
```

# Model Spec
```rb
RSpec.describe User, type: :model do
  it 'has a valid factory' do # 有効なファクトリを持つこと
    expect(build(:user).valid?).to eq true
  end
end
```
