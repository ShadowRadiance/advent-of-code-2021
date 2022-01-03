# frozen_string_literal: true

require "./infinite_image"

RSpec.describe InfiniteImage do
  subject { InfiniteImage.new(data, background: background) }

  let(:data) {
    [
      "#..#.",
      "#....",
      "##..#",
      "..#..",
      "..###",
    ]
  }

  context "with a dark background" do
    let(:background) { "." }
    
    it "can tell it's size" do
      expect([
        subject.min_x, 
        subject.min_y, 
        subject.max_x, 
        subject.max_y, 
        subject.width, 
        subject.height
      ]).to eq([
        0,0,4,4,5,5
      ])
    end
  
    it "correctly reports the lit lights" do
      expect(subject.count_lit).to eq(10)
    end
  
    it "correctly reports the unlit lights" do
      expect(subject.count_unlit).to eq(Float::INFINITY)
    end
  
    it "correctly reports each light state" do
      expect(subject.nine_lights_around(-1, -1)).to eq([
        [".", ".", "."],
        [".", ".", "."],
        [".", ".", "#"],
      ])
  
      expect(subject.nine_lights_around(1, 1)).to eq([
        ["#", ".", "."],
        ["#", ".", "."],
        ["#", "#", "."],
      ])
  
      expect(subject.nine_lights_around(4, 3)).to eq([
        [".", "#", "."],
        [".", ".", "."],
        ["#", "#", "."],
      ])
    end
  end

  context "with a light background" do
    let(:background) { "#" }

    it "can tell it's size" do
      expect([
        subject.min_x, 
        subject.min_y, 
        subject.max_x, 
        subject.max_y, 
        subject.width, 
        subject.height
      ]).to eq([
        0,0,4,4,5,5
      ])
    end
  
    it "correctly reports the lit lights" do
      expect(subject.count_lit).to eq(Float::INFINITY)
    end
  
    it "correctly reports the unlit lights" do
      expect(subject.count_unlit).to eq(15)
    end

    it "correctly reports each light state" do
      expect(subject.nine_lights_around(-1, -1)).to eq([
        ["#", "#", "#"],
        ["#", "#", "#"],
        ["#", "#", "#"],
      ])
  
      expect(subject.nine_lights_around(1, 1)).to eq([
        ["#", ".", "."],
        ["#", ".", "."],
        ["#", "#", "."],
      ])
  
      expect(subject.nine_lights_around(4, 3)).to eq([
        [".", "#", "#"],
        [".", ".", "#"],
        ["#", "#", "#"],
      ])
    end
  end
end
