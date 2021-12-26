# frozen_string_literal: true

require "./image_enhancer"

RSpec.describe ImageEnhancer do
  subject { ImageEnhancer.new(algorithm) }
  let(:algorithm) { "..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..###..######.###...####..#..#####..##..#.#####...##.#.#..#.##..#.#......#.###.######.###.####...#.##.##..#..#..#####.....#.#....###..#.##......#.....#..#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#......#..#...##.#.##..#...##.#.##..###.#......#.#.......#.#.#.####.###.##...#.....####.#..#..#.##.#....##..#.####....##...##..#...#......#.#.......#.......##..####..#...#.#.#...##..#.#..###..#####........#..####......#..#" }
  let(:image) { InfiniteImage.new(data) }
  let(:data) {
    [
      "#..#.",
      "#....",
      "##..#",
      "..#..",
      "..###",
    ]
  }

  context "after 0 enhancements" do
    it "has the correct number of lights" do
      expect(image.count_lit).to eq(10)
    end

    it "can display the enhanced image" do
      expect(image.to_s).to eq(
        <<~IMAGE.chomp
        #..#.
        #....
        ##..#
        ..#..
        ..###
        IMAGE
      )
    end
  end

  context "after 1 enhancement" do
    before { @result = subject.enhance(image) }

    it "has the correct number of lights" do
      expect(@result.count_lit).to eq(24)
    end

    it "can display the enhanced image" do
      expect(@result.to_s).to eq(
        <<~IMAGE.chomp
        .##.##.
        #..#.#.
        ##.#..#
        ####..#
        .#..##.
        ..##..#
        ...#.#.
        IMAGE
      )
    end
  end
  
  context "after 2 enhancements" do
    before { @result = subject.enhance(subject.enhance(image)) }
    
    it "has the correct number of lights" do
      expect(@result.count_lit).to eq(35)
    end
    
    it "can display the enhanced image" do
      expect(@result.to_s).to eq(
        <<~IMAGE.chomp
        .......#.
        .#..#.#..
        #.#...###
        #...##.#.
        #.....#.#
        .#.#####.
        ..#.#####
        ...##.##.
        ....###..
        IMAGE
      )
    end

  end
end
