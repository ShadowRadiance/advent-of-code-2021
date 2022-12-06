require "./origami_parser"

RSpec.describe OrigamiParser do
  subject { described_class.new(data) }

  let(:data) {
    <<~DATA
      6,10
      0,14
      9,10
      0,3
      10,4
      4,11
      6,0
      6,12
      4,1
      0,13
      10,12
      3,4
      3,0
      8,4
      1,10
      2,14
      8,10
      9,0

      fold along y=7
      fold along x=5
    DATA
  }

  it "extracts the dots" do
    expect(subject.dots.map{ |dot| "#{dot.x} / #{dot.y}" }).to eq([
      "6 / 10",
      "0 / 14",
      "9 / 10",
      "0 / 3",
      "10 / 4",
      "4 / 11",
      "6 / 0",
      "6 / 12",
      "4 / 1",
      "0 / 13",
      "10 / 12",
      "3 / 4",
      "3 / 0",
      "8 / 4",
      "1 / 10",
      "2 / 14",
      "8 / 10",
      "9 / 0",
    ])
  end

  it "extracts the instructions" do
    expect(subject.instructions.map { |inst| "#{inst.axis} @ #{inst.index}" }).to eq([
      "y @ 7",
      "x @ 5",
    ])
  end
end