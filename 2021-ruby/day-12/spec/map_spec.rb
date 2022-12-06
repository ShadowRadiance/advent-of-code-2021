require "./cave/map"

RSpec.describe Cave::Map do
  subject { described_class.new(input) }

  let(:input) {
    double(
      nodes: Set.new(%w[start A b c d end]),
      connections: [
        double(from: "start", to: "A"),
        double(from: "start", to: "b"),
        double(from: "A", to: "c"),
        double(from: "A", to: "b"),
        double(from: "b", to: "d"),
        double(from: "A", to: "end"),
        double(from: "b", to: "end"),
      ]
    )
  }

  it "graph" do
    expect(subject.rooms.count).to eq(6)
    expect(subject.rooms.select(&:big?).count).to eq(1)
    expect(subject.rooms.select(&:small?).count).to eq(5)
    expect(subject.find_room("start").connected_rooms.map(&:name)).to eq(["A", "b"])
  end
end