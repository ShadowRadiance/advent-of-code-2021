require "./cave/input"
require 'set'

RSpec.describe Cave::Input do
  subject { described_class.new(data) }

  let(:data) {
    <<~DATA
      start-A
      start-b
      A-c
      A-b
      b-d
      A-end
      b-end
    DATA
  }

  it "extracts the unique nodes" do
    expect(subject.nodes).to eq(Set.new(%w[start A b c d end]))
  end

  it "extracts the unique nodes" do
    expect(subject.connections).to eq([
      Cave::Input::Connection.new("start", "A"),
      Cave::Input::Connection.new("start", "b"),
      Cave::Input::Connection.new("A", "c"),
      Cave::Input::Connection.new("A", "b"),
      Cave::Input::Connection.new("b", "d"),
      Cave::Input::Connection.new("A", "end"),
      Cave::Input::Connection.new("b", "end"),
    ])
  end
end